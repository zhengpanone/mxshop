# 构建阶段
FROM python:3.9 as builder

WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt


# 运行阶段
FROM python:3.9-slim

RUN useradd --create-home --shell /bin/bash app

# 复制全量依赖
COPY --from=builder /usr/local /usr/local

USER app
WORKDIR /home/app


ENV PATH=/usr/local/bin:$PATH
ENV PYTHONPATH=/usr/local/lib/python3.9/site-packages

COPY --chown=app:app . .

WORKDIR /home/app/userop_srv

EXPOSE 5001

CMD ["python", "-m", "server", "--host", "0.0.0.0", "--port", "5001"]
