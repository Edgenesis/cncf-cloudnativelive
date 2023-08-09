# file: dockerfile
FROM python:3-slim

COPY driver.py .
COPY requirements.txt .

RUN apt-get update && \
    apt-get install -yq \
    gcc && \
    pip3 install --no-cache-dir -r requirements.txt && \
    apt-get clean && \
    apt-get remove -yq gcc && \
    rm -rf /var/lib/apt/lists/*

ENTRYPOINT [ "python3", "driver.py" ]