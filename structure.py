import os

def generate_markdown(directory, depth=0):
    """
    Recursively generates markdown list representing project directory structure.
    """
    markdown = ""
    indent = "    " * depth  # Indentation for the current level

    # List all files and directories in the current directory
    for item in sorted(os.listdir(directory)):
        item_path = os.path.join(directory, item)

        if os.path.isdir(item_path):  # If it's a directory, recurse into it
            markdown += f"{indent}- **{item}**\n"
            markdown += generate_markdown(item_path, depth + 1)  # Recurse into subdirectories
        else:  # If it's a file, add it to the markdown
            markdown += f"{indent}- {item}\n"

    return markdown

def create_project_structure_markdown(root_directory, output_file):
    """
    Creates a markdown file representing the project structure.
    """
    project_structure = generate_markdown(root_directory)

    with open(output_file, "w", encoding="utf-8") as f:
        f.write("# Project Directory Structure\n\n")
        f.write(project_structure)

if __name__ == "__main__":
    root_dir = "/Users/yuntwo/Projects/Go/finTech/app/coupon"  # Change this to the path of your project directory
    output_file = "README-structure-temp.md"  # Output file for the markdown
    create_project_structure_markdown(root_dir, output_file)
    print(f"Project structure has been written to {output_file}")
