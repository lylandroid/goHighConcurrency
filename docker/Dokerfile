FROM frolvlad/alpine-oraclejdk8:slim
VOLUME /tmp
ADD dockerdemo-0.0.1-SNAPSHOT.jar app.jar
ENTRYPOINT ["java","-jar","/app.jar"]