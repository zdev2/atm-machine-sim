# ATM Machine

This is a simple ATM Machine simulation written in Go. It allows users to perform basic banking operations such as checking balance, withdrawing cash, and depositing cash using a simulated ATM card.

## Features

- Insert an ATM card and login with a password.
- Check balance.
- Withdraw cash.
- Deposit cash.
- User-friendly command-line interface.

## Getting Started

### Prerequisites

- Go 1.16 or higher.

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/zdev2/atm-machine-sim.git
   cd atm-machine-sim
   ```

2. Build the project:

   ```sh
   go build -o atm-machine-sim.exe main.go
   ```

3. The repository includes a sample card located at `card-folder`:

   ```json
   {
     "card-num": "1234-5678-9012-3456",
     "card-user": "John Doe",
     "card-pass": "password123",
     "card-balance": 1000000
   }
   ```

### Usage

1. Run the program:

   ```sh
   ./atm-machine-sim
   ```

2. Follow the on-screen instructions to insert the card, login, and perform various banking operations.

## Project Structure

- `main.go`: The main file containing the program logic.
- `card-slot/card.json`: The simulated ATM card data.
- `helpers`: A package containing helper functions used in the program.

### Contributing

Contributions are welcome! Please open an issue or submit a pull request.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Acknowledgements

- [Go Programming Language](https://golang.org/)
- [JSON](https://www.json.org/)
