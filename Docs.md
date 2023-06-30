# Text-to-Speech API

The Text-to-Speech API allows you to convert text into speech using various voice models. It provides both GET and POST endpoints for generating audio output.

## Base URL

```
http://localhost:8080/
```

## Endpoints

### Convert Text to Speech (POST)

Converts the provided text into speech using the specified voice model.

```
POST /api/tts
```

#### Request Body

The request body should be a JSON object with the following properties:

- `text` (string, required): The text to be converted into speech.
- `voice` (string, optional): The voice model to be used for speech synthesis. If not provided, the default voice model will be used.

##### Example Request Body

```json
{
  "text": "Hello, world!",
  "voice": "en-us-libritts-low.onnx"
}
```

#### Response

- Status Code: 200 (OK)
- Content-Type: audio/wav

The response will be an audio file in WAV format containing the synthesized speech.

##### Example Response

```
HTTP/1.1 200 OK
Content-Type: audio/wav
Content-Length: <file_size>

<binary audio data>
```

### Convert Text to Speech (GET)

Converts the provided text into speech using the specified voice model.

```
GET /api/tts?text=<text>&voice=<voice>
```

#### Parameters

- `text` (string, required): The text to be converted into speech.
- `voice` (string, optional): The voice model to be used for speech synthesis. If not provided, the default voice model will be used.

#### Response

- Status Code: 200 (OK)
- Content-Type: audio/wav

The response will be an audio file in WAV format containing the synthesized speech.

##### Example Response

```
HTTP/1.1 200 OK
Content-Type: audio/wav
Content-Length: <file_size>

<binary audio data>
```

### Get List of Available Voices

Retrieves a list of available voice models that can be used for speech synthesis.

```
GET /api/voices
```

#### Response

- Status Code: 200 (OK)
- Content-Type: application/json

The response will be a JSON array containing the names of available voice models.

##### Example Response

```
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: <response_size>

[
  "en-us-libritts-high.onnx",
  "en-us-libritts-low.onnx",
  ...
]
```

## Error Responses

In case of errors, the API will return an appropriate HTTP status code along with an error message in the response body.

- Status Code: 400 (Bad Request)
- Content-Type: text/plain

```
HTTP/1.1 400 Bad Request
Content-Type: text/plain

Error parsing JSON: Invalid request body
```

- Status Code: 500 (Internal Server Error)
- Content-Type: text/plain

```
HTTP/1.1 500 Internal Server Error
Content-Type: text/plain

Error running executable: Failed to start speech synthesis process
```
