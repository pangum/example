FROM storezhang/alpine


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL description="子云视频转码平台执行器"


# 复制文件
COPY docker /
COPY agent /opt/ziyunix/


RUN set -ex \
    \
    \
    \
    && apk update \
    && apk add --no-cache ffmpeg \
    \
    \
    \
    # 增加执行器的执行权限
    && chmod +x /etc/s6/ziyunix/* \
    && chmod +x /usr/bin/ziyunix \
    \
    \
    \
    # 清理镜像，减少无用包
    && rm -rf /var/cache/apk/*


# 设置子云主目录
ENV ZIYUNIX_HOME /config
