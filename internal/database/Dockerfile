FROM postgres:16

# Set the timezone
ENV TZ=Europe/Bucharest
RUN ln -fs /usr/share/zoneinfo/$TZ /etc/localtime && dpkg-reconfigure -f noninteractive tzdata

EXPOSE 5432
# Expose PostgreSQL port

