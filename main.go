package main

import (
    "golang.org/x/crypto/ssh"
    "os"
    "log"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
    
    // Cargar la llave privada
    privateKeyPath := os.Getenv("PRIVATE_KEY_PATH")
    privateKey, err := os.ReadFile(privateKeyPath)
    if err != nil {
        log.Fatalf("unable to read private key: %v", err)
    }

    // Crear el Signer a partir de la llave privada
    signer, err := ssh.ParsePrivateKey(privateKey)
    if err != nil {
        log.Fatalf("unable to parse private key: %v", err)
    }

    // Configuración del cliente SSH
    config := &ssh.ClientConfig{
        User: os.Getenv("SSH_USERNAME"), // Reemplaza esto con tu nombre de usuario en la instancia GCP
        Auth: []ssh.AuthMethod{
            ssh.PublicKeys(signer),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(), // No es recomendado para producción
    }

    // Dirección del servidor
    host := os.Getenv("SSH_IP_TEST1") // Reemplaza "server_ip" con la IP del servidor
    client, err := ssh.Dial("tcp", host, config)
    if err != nil {
        log.Fatalf("Failed to dial: %v", err)
    }

    // Abre una sesión
    session, err := client.NewSession()
    if err != nil {
        client.Close()
        log.Fatalf("Failed to create session: %v", err)
    }
    defer session.Close()

    // Ejecuta un comando
    output, err := session.CombinedOutput("ls -l") // Cambia "ls -l" por el comando que necesites ejecutar
    if err != nil {
        log.Fatalf("Failed to run command: %v", err)
    }
    log.Println(string(output))
}
