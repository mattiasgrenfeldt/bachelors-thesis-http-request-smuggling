FROM python:3.9.1
RUN pip install gunicorn==20.0.4
WORKDIR /app
COPY app.py .
CMD ["gunicorn", "-b", "0.0.0.0:80", "app:app"]
