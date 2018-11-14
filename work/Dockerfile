FROM python:3.7-slim

WORKDIR /app
COPY . /app
RUN pip install -r requirements.txt

ENTRYPOINT ["python", "work.py"]
