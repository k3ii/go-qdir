
# Qdir - Quick Directory Generator

Qdir (Quick Directory) is a versatile command-line tool for creating directories with customizable naming schemes. It streamlines the process of creating uniquely named directories for various purposes such as temporary workspaces, project organization, or test environments.

## Features

- Generate directories with random hexadecimal names or names of notable scientists/technologists
- Create nested directory structures with controllable depth
- Option to use system's temporary directory or current working directory
- Adjustable length for hexadecimal names

## Installation

To install Qdir, make sure you have Go installed on your system, then run:

```bash
go get -u github.com/k3ii/qdir
```

Replace `yourusername` with your actual GitHub username or the appropriate path to the repository.

## Usage

Basic usage:

```bash
qdir [flags]
```

### Flags

- `-n, --nested <depth>`: Depth of nested directories to create (default 0)
- `-u, --use-names`: Use scientist/technologist names instead of random hex
- `-t, --tmp`: Use the system's temporary directory
- `-l, --hex-length <length>`: Length of the random hexadecimal name (default 16)

### Examples

1. Create a single directory with a random hex name in the current directory:
   ```bash
   qdir
   ```

2. Create a nested directory structure 3 levels deep with scientist names in the temp directory:
   ```bash
   qdir -n 3 -u -t
   ```

3. Create a directory with a 32-character hex name:
   ```bash
   qdir -l 32
   ```

## Contributing

Contributions to Qdir are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by Docker's naming generator for containers
- Thanks to all the scientists and technologists whose names are used in this project

## Support

If you encounter any problems or have any questions, please open an issue on the GitHub repository.
