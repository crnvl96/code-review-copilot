#!/bin/bash

echo "Starting Ollama server..."
ollama serve &

echo "Waiting for Ollama server to be active..."
while [ "$(ollama list | grep 'NAME')" == "" ]; do
  sleep 1
done

echo "Ollama server active, pulling Tinyllama..."
ollama pull tinyllama

echo "Tinyllama is ready to be used!!!"
