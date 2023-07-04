# serve-piper-tts

Voice Samples from Piper - https://rhasspy.github.io/piper-samples/

![image](https://github.com/arunk140/serve-piper-tts/assets/8670239/b0c7e4cf-d3cd-42f1-bb29-46674a762e45)

Hosted API Test Server (Temporary Link) - https://voice.arunk140.com/ | https://voice.arunk140.com/api/tts?text=github

Go Lang API Wrapper around Piper TTS - Supports TTS Inference and List of Voices

---

### Install with Docker

```
docker pull ghcr.io/arunk140/serve-piper-tts:latest
docker run -p 8080:8080 ghcr.io/arunk140/serve-piper-tts:latest
```

OR

```
docker compose up -d
```

---

Add you Piper Voice Models from or use the download script (to the models directory)

- https://github.com/rhasspy/piper/releases/tag/v0.0.2
- https://huggingface.co/rhasspy/piper-voices/tree/main
- ./download-voices.sh

To download and extract specific files for a language, use the following format:

./download-voices.sh LANG_CODE

e.g. to download en (English) Voices

```
chmod +x ./download-voices.sh
./download-voices.sh en
```

Check the download-voices.sh file for a list of voices and supported languages.

Get the Latest Piper Executable from Piper's GitHub Releases or using the Download Script (download and extract in the same folder)

- https://github.com/rhasspy/piper/releases/latest
- ./download-piper.sh

```
chmod +x ./download-piper.sh
./download-piper.sh
```

To run the API server directy -

```
go run .
```

To Build executable and Run -

```
go build
./serve-piper-go
```

Runs on Port 8080 By default - Update main.go constant to change port

```
http://localhost:8080/
```

API Docs in [Docs.md](Docs.md)
