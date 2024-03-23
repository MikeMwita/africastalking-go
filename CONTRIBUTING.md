
# Contributing to African's Talking Go SDK

We welcome contributions from the community! If you'd like to contribute to the African's Talking Go SDK, please follow these guidelines.

## Setting Up Your Environment

Before you can contribute, you need to set up your development environment. Here's how you can do it:

1. **Fork the Repository**
    - Go to the [African's Talking Go SDK repository](https://github.com/MikeMwita/africastalking-go).
    - Click on the "Fork" button in the top-right corner of the page and clone your fork.

    
3. **Install Dependencies**
    - Navigate to the cloned repository directory in your terminal.
    - Run the following command to install the necessary dependencies:
      ```bash
      go mod tidy
      ```

4. **Create a New Branch**
    - Create a new branch for your feature or bug fix:
      ```bash
      git checkout -b your-branch-name
      ```
    - Replace `your-branch-name` with a descriptive name for your branch.

## Making Changes

1. **Make Your Changes**
    - Open the project in your favorite IDE or editor.
    - Make your changes to the codebase.

2. **Write Tests**
    - Add tests that cover the changes you've made.
    - Ensure all tests pass by running:
      ```bash
      go test -cover ./...
      ```

3. **Document Your Changes**
    - Update the `README.md` if necessary.
    - Add comments to your code where appropriate.

4. **Commit Your Changes**
    - Stage your changes for commit:
      ```bash
      git add .
      ```
    - Commit your changes with a descriptive message:
      ```bash
      git commit -m "Add a brief description of your changes"
      ```

5. **Push to Your Fork**
    - Push your changes to your fork on GitHub:
      ```bash
      git push origin your-branch-name
      ```

## Submitting a Pull Request

1. **Open a Pull Request**
    - Click on "Pull requests" and then the "New pull request" button.
    - Choose your fork and branch as the source and the original repository's main branch as the target.
    - Fill in the pull request template with information about your changes.

2. **Code Review**
    - Wait for the maintainers to review your pull request.
    - Make any requested changes.

3. **Merge**
    - Once your pull request is approved, the maintainers will merge it into the main codebase.

Thank you for contributing to the African's Talking Go SDK!

---

 Happy coding! 🚀