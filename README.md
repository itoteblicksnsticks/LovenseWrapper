# LovenseWrapper

1. Obtain your Lovense API access code. If you don't have one, please contact Lovense support to request access

2. Clone this repository to your machine

3. Open the `main.go` file and replace `"your_access_code_here"` with your actual Lovense API access code

4. Build the project by running the following command:
   ```
   go build
   ```

## Features

- Retrieve a list of connected Lovense toys
- Control your toys with various commands like vibration, rotation, pumping, and more
- Supports local server mode and command-line mode
- Fine-grained control over strength, duration, rotation, air level, and vibration patterns
- Option to loop commands for continuous stimulation

## Usage

LovenseWrapper provides two ways to interact with your Lovense toys:

1. **Local Server Mode**:
   - Run the executable with the `-local` flag to start a local server at `http://localhost:8080`
   - Send HTTP requests to the `/toys` endpoint to retrieve the list of connected toys
   - Send HTTP POST requests to the `/control` endpoint with a JSON payload to control your toys

2. **Command-line Mode**:
   - Run the executable without any flags to enter the command-line mode
   - Follow the prompts to enter the toy ID, command, and other parameters
   - The command will be sent to the specified toy, and the result will be displayed in the console

## Examples

Here are a few examples of how to use LovenseWrapper:

## Vibrate your toy at strength 10 for 5 seconds
```
Enter Toy ID: toy_123
Enter Command: Vibrate
Enter Strength (0-20): 10
Enter Vibration Pattern (0-3): 0
Enter Duration (in seconds): 5
Loop Command? (true/false): false
```

### Rotate your toy at strength 15 for 10 seconds in a loop
```
Enter Toy ID: toy_456
Enter Command: Rotate
Enter Strength (0-20): 15
Enter Duration (in seconds): 10
Loop Command? (true/false): true
```

## Contributing

Contributions to LovenseWrapper are welcome, If you find any bugs, have suggestions for improvements, or want to add new features, please feel free to open an issue or submit a pull request. Make sure to follow the existing code style and include appropriate documentation and tests for your changes


Example POST request body to the `/control` endpoint:

```json
{
  "toyID": "toy_123",
  "command": "Vibrate",
  "strength": 10,
  "duration": 5,
  "loop": false,
  "rotation": 0,
  "pump": 0,
  "airLevel": 0,
  "vibration": 0
}
```

In this example, the POST request body is a JSON object containing the following fields:
- `toyID`: The ID of the toy to control
- `command`: The command to send to the toy (e.g, "Vibrate", "Rotate", "Pump", etc)
- `strength`: The strength of the vibration or rotation (0-20)
- `duration`: The duration of the command in seconds
- `loop`: Whether to loop the command (true/false)
- `rotation`: The rotation pattern (0-3)
- `pump`: The pump status (0-3)
- `airLevel`: The air level for AirIn/AirOut commands (0-3)
- `vibration`: The vibration pattern (0-3)

Make sure to adjust the values according to your specific requirements when sending the POST request
