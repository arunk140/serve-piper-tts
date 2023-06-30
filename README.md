# serve-piper-tts
Go Lang API Wrapper around Piper TTS - Supports TTS Inference and List of Voices

Add you Piper Voice Models from or use the download script (to the models directory)

- https://github.com/rhasspy/piper/releases/tag/v0.0.2
- https://huggingface.co/rhasspy/piper-voices/tree/main
- ./download-voices.sh


Get the Latest Piper Executable from Piper's GitHub Releases or using the Download Script (download and extract in the same folder)

- https://github.com/rhasspy/piper/releases/latest
- ./download-piper.sh


To run the API server directy - 
```
go run .
```

To Build executable and Run - 
```
go build
./serve-piper-go
```

API Docs in *[Docs.md](Docs.md)*
