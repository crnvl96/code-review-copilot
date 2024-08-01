FROM ollama/ollama:latest

WORKDIR /app

COPY . ./root/.ollama/

COPY ./run-ollama.sh ./tmp/run-ollama.sh

RUN chmod +x ./tmp/run-ollama.sh \
    && ./tmp/run-ollama.sh

EXPOSE 11434
