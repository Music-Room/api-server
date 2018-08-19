FROM scratch

ADD m usic-room /

CMD mkdir config

#COPY config/music-room.yaml config/

ENTRYPOINT ["/music-room"]