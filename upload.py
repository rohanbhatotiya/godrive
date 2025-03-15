import sys
import zipfile
import os
import io
import shutil
from tqdm import tqdm
from pydrive2.auth import GoogleAuth
from pydrive2.drive import GoogleDrive

# Set up the directory to store authentication and secrets
godrive_dir = os.path.expanduser("~/.godrive")
os.makedirs(godrive_dir, exist_ok=True)

auth_file_path = os.path.join(godrive_dir, "auth.txt")
client_secrets_path = os.path.join(godrive_dir, "client_secrets.json")

def ensure_client_secrets():
    """Ensures `client_secrets.json` is present, otherwise prompts the user to provide it."""
    if not os.path.exists(client_secrets_path):
        print("Missing or invalid client_secrets.json. You need to provide a new one.")
        print("\nPaste your client_secrets.json content below (Copy-paste the entire JSON, then press Enter twice):\n")

        user_input = []
        while True:
            try:
                line = input()
                if not line.strip():  # Detect double enter (empty line)
                    break
                user_input.append(line)
            except KeyboardInterrupt:
                print("\nOperation cancelled. Exiting safely.")
                sys.exit(1)

        client_json_content = "\n".join(user_input)

        try:
            with open(client_secrets_path, "w") as f:
                f.write(client_json_content)
            print(" Successfully verified and saved `client_secrets.json` at ~/.godrive/")
        except Exception as e:
            print(f"Failed to save `client_secrets.json`: {e}")
            sys.exit(1)


def authenticate():
    """Handles Google authentication, prompting the user if needed."""
    ensure_client_secrets()

    gauth = GoogleAuth()
    gauth.LoadClientConfigFile(client_secrets_path)

    if os.path.exists(auth_file_path):
        gauth.LoadCredentialsFile(auth_file_path)
        if gauth.credentials and not gauth.access_token_expired:
            print(" Using existing authentication.")
            return GoogleDrive(gauth)

    print(" Re-authenticating...")
    gauth.CommandLineAuth()
    gauth.SaveCredentialsFile(auth_file_path)
    print(" Authentication successful.")

    return GoogleDrive(gauth)


def zip_directory(directory_path):
    """Creates a zip file for a directory before uploading."""
    parent_dir = os.path.dirname(directory_path)
    dir_name = os.path.basename(directory_path)
    zip_name = os.path.join(parent_dir, f"{dir_name}(by_godrive).zip")

    if os.path.exists(zip_name):
        print(f" Using existing zip: {zip_name}")
        return zip_name

    print(" Creating a zip file...")
    try:
        with zipfile.ZipFile(zip_name, 'w', zipfile.ZIP_DEFLATED) as zipf:
            for root, _, files in os.walk(directory_path):
                for file in files:
                    file_path = os.path.join(root, file)
                    zipf.write(file_path, os.path.relpath(file_path, directory_path))
        print(f" Zip created: {zip_name}")
    except Exception as e:
        if "No space left" in str(e):
            print(" No space left to create a zip. Try uploading files individually.")
            sys.exit(1)
        else:
            print(f"Error while zipping: {e}")
            sys.exit(1)

    return zip_name


def upload_file(drive, file_path, custom_name):
    """Uploads a file or directory (as zip) to Google Drive."""
    if not os.path.exists(file_path):
        print(" Error: File does not exist!")
        return False

    is_zipped = False

    if os.path.isdir(file_path):
        file_path = zip_directory(file_path)
        is_zipped = True

    # Extract original name and extension
    original_name, original_extension = os.path.splitext(os.path.basename(file_path))

    if custom_name:
        upload_name = f"{custom_name}{original_extension}"
    else:
        upload_name = f"{original_name}{original_extension}"

    file_drive = drive.CreateFile({"title": upload_name})

    print("\n Uploading file...")
    try:
        with open(file_path, "rb") as f:
            file_drive.content = io.BytesIO(f.read())
            file_drive.Upload()

        print("\nUpload Complete!")
        print(f" File uploaded with name: {upload_name}")

        if is_zipped:
            os.remove(file_path)

    except Exception as e:
        print(f" Upload failed: {e}")
        return False

    return True


def main():
    """Handles command-line arguments and uploads a file/directory."""
    if len(sys.argv) < 2:
        print(" Error: No file path provided. Exiting...")
        sys.exit(1)

    file_path = sys.argv[1]
    custom_name = sys.argv[2] if len(sys.argv) > 2 else ""

    print(f" Uploading: {file_path}")
    print(f"\tUpload Name: {custom_name if custom_name else 'Using original name'}")

    drive = authenticate()
    success = upload_file(drive, file_path, custom_name)
    if not success:
        print(" Upload failed. Please try again.")
        sys.exit(1)


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print(" Operation cancelled by user. Exiting safely...")
        sys.exit(1)
