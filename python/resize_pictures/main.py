import os
from PIL import Image, ImageFilter
from pathlib import Path

allowed_files_format = ["jpg", "png", "jpeg", "jpe"]

pdf_compression_level = 9
target_width = 1080

#
#   TOOLS FUNCTIONS
#


def rotate_image(image: Image) -> Image:
    rotated_image = image
    # Rotate the image based on EXIF orientation
    exif = image.getexif()
    if exif:
        for tag, orientation in exif.items():
            if tag == 274:
                if orientation == 1:
                    rotated_image = image.rotate(-90, expand=True)
                elif orientation == 3:
                    rotated_image = image.rotate(270, expand=True)
                elif orientation == 6:
                    rotated_image = image.rotate(270, expand=True)
                elif orientation == 8:
                    rotated_image = image.rotate(90, expand=True)
                break  # Exit the loop once orientation tag is found
    return rotated_image


def calculate_new_picture_size(width: int, height) -> tuple[int, int]:
    aspect_ratio = width / height
    new_width = target_width
    new_height = int(new_width / aspect_ratio)
    return (new_width, new_height)


def reduce_image_size(input_path, output_path):
    # Open file
    img = Image.open(input_path)
    # Get new height and width
    width, height = img.size
    new_width, new_height = calculate_new_picture_size(width, height)
    # Resize the image
    resized_img = img.resize((new_width, new_height), Image.LANCZOS)
    # Rotate the image
    rotated_img = rotate_image(resized_img)
    # Apply sharpening filter
    sharpened_img = rotated_img.filter(ImageFilter.SHARPEN)
    # Save picture
    sharpened_img.save(output_path, optimize=True, quality=100)
    # Close the images
    img.close()
    resized_img.close()


def get_non_py_files(folder_path) -> list[str]:
    non_py_files: list[str] = []
    for filename in os.listdir(folder_path):
        ext = os.path.splitext(filename)[1][1:].lower()
        if ext in allowed_files_format:
            non_py_files.append(str(filename))
    return non_py_files


#
#   STEPS FUNCTION
#


def execute_reduce_picture_size_steps():
    # folder = input("Enter the source folder path: ")
    folder = "/Users/medici/Desktop/ai/Pictures/recipes"
    origin_folder_path = Path(folder)
    target_folder_path = origin_folder_path / "resized"
    # Create the target folder if it doesn't exist
    target_folder_path.mkdir(parents=True, exist_ok=True)
    # Get the filenames
    filenames = get_non_py_files(origin_folder_path)
    count = 0
    for filename in filenames:
        count += 1
        print(f"Reducing size of {filename} ({count}/{len(filenames)})")
        # reduce_image_size(origin_folder_path / filename, target_folder_path / filename)


if __name__ == "__main__":
    execute_reduce_picture_size_steps()
