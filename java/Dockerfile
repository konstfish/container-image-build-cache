ARG BUILD_IMAGE=maven
ARG BUILD_TAG=3.9.7-eclipse-temurin-21

FROM ${BUILD_IMAGE}:${BUILD_TAG} as builder

COPY pom.xml .
COPY src src

RUN mvn clean package

FROM eclipse-temurin:21.0.2_13-jdk

COPY --from=builder /target/*.jar app.jar

EXPOSE 8080

ENTRYPOINT ["java", "-jar", "/app.jar"]