# Go Project Setup Guide

This guide outlines the steps to set up and run the Go project.

## Prerequisites

Ensure you have the following installed:

- Go (latest version)
- Node.js and npm (latest version)
- Prisma CLI (installed via npm or npx)

---

## Steps to Set Up the Project

1. **Copy `.env.example` to `.env`**

   ```bash
   cp .env.example .env
   ```

2. **Install Node.js Dependencies**

   ```bash
   npm install
   ```

3. **Install Prisma Client for Go**

   ```bash
   go get github.com/steebchen/prisma-client-go
   ```

4. **Tidy Go Modules**

   ```bash
   go mod tidy
   ```

5. **Run Prisma Migrations**

   ```bash
   npx prisma migrate
   ```

6. **Generate Prisma Client**

   ```bash
   npx prisma generate
   ```

7. **Run the Application**

   ```bash
   go run .
   ```

---

## Notes

- Ensure the `.env` file contains the correct configuration values for your database and environment.
- Use `npx prisma studio` to inspect your database visually if needed.
- Check `go.mod` and `go.sum` to confirm dependencies are installed properly.

Follow these steps, and your Go project should be up and running smoothly!
