# debug-on-cursor-101

![ScreenShot](Screenshot.png)

## Environment Variables Management (`.env`)

- Store your environment variables in a `.env` file at the project root. Example:
  ```env
  PORT=8080
  DB_CONN=postgres://user:pass@localhost:5432/db
  SECRET_KEY=your-secret-key
  ```
- The application reads these variables using `os.Getenv` in `main.go`.
- **Note:** The `.env` file is loaded automatically when debugging in VS Code (see below).

## Running the Application

- **Standard Run:**
  ```sh
  export $(grep -v '^#' .env | xargs) # Load .env variables into your shell
  go run main.go
  ```
  Or set variables inline:
  ```sh
  PORT=8080 DB_CONN=... SECRET_KEY=... go run main.go
  ```

## Debugging in VS Code

- Open the project in VS Code.
- Press `F5` or go to the Run & Debug panel and select **Debug Go Program with Env**.
- This will launch the app with variables from `.env` automatically loaded (see `.vscode/launch.json`).
- You can set breakpoints in `main.go` (e.g., inside the `add` function or anywhere else).

---

For more details, see `.vscode/launch.json` for debug configuration and `main.go` for environment variable usage.
