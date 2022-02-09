pipeline {
  environment {
    networkName = "network-${env.BUILD_TAG}"
    registry = 'tyrm/gemini-forum'
    registryCredential = 'docker-io-tyrm'
    dockerImage = ''
    gitDescribe = ''
  }

  agent any

  stages {

    stage('Setup') {
      steps {
        script {
          gitDescribe = sh(returnStdout: true, script: 'git describe --tag').trim()
          writeFile file: "./version/version.go", text: """package version

// Version of the application
const Version = "${gitDescribe}"

          """
        }
      }
    }

    stage('Start Postgres'){
      steps{
        script{
          retry(4) {
            echo 'trying to start postgres'
            withCredentials([
              usernamePassword(credentialsId: 'integration-postgres-test', usernameVariable: 'POSTGRES_USER', passwordVariable: 'POSTGRES_PASSWORD'),
              string(credentialsId: 'integration-redis-test', variable: 'REDIS_PASSWORD')
            ]) {
              sh """NETWORK_NAME="${networkName}" docker-compose -f docker-compose-integration.yaml pull
              NETWORK_NAME="${networkName}" docker-compose -p ${env.BUILD_TAG} -f docker-compose-integration.yaml up -d"""
            }
          }
        }
      }
    }

    stage('Test') {
      agent {
        docker {
          image 'gobuild:1.17'
          args '--network ${networkName} -e GOCACHE=/gocache -e HOME=${WORKSPACE} -v /var/lib/jenkins/gocache:/gocache'
        }
      }
      steps {
        script {
          withCredentials([
            string(credentialsId: 'codecov-tyrm-gemini-forum', variable: 'CODECOV_TOKEN'),
            usernamePassword(credentialsId: 'integration-postgres-test', usernameVariable: 'POSTGRES_USER', passwordVariable: 'POSTGRES_PASSWORD'),
            string(credentialsId: 'integration-redis-test', variable: 'REDIS_PASSWORD')
          ]) {
            pgConnectionDSN = "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/supremerobot?sslmode=disable"

            sh """#!/bin/bash
            go get -t -v ./...
            TEST_DSN="${pgConnectionDSN}" TEST_REDIS="redis:6379" TEST_REDIS_PASS="${REDIS_PASSWORD}" go test --tags=integration -race -coverprofile=coverage.txt -covermode=atomic ./...
            RESULT=\$?
            gosec -fmt=junit-xml -out=gosec.xml  ./...
            bash <(curl -s https://codecov.io/bash)
            exit \$RESULT
            """
          }
          junit allowEmptyResults: true, checksName: 'Security', testResults: "gosec.xml"
        }
      }
    }

    stage('Upload image') {
      steps {
        script {
          retry(3) {
            timeout(time: 15, unit: 'MINUTES') {
              if (env.TAG_NAME) {
                sh "DOCKER_CLI_EXPERIMENTAL=enabled docker buildx build --platform linux/arm64,linux/amd64 -t ${registry}:${env.TAG_NAME} . --push"
              } else {
                sh "DOCKER_CLI_EXPERIMENTAL=enabled docker buildx build --platform linux/arm64,linux/amd64 -t ${registry}:${env.BRANCH_NAME} . --push"
              }
            }
          }
        }
      }
    }

  }

  post {
    always {
      withCredentials([
        usernamePassword(credentialsId: 'integration-postgres-test', usernameVariable: 'POSTGRES_USER', passwordVariable: 'POSTGRES_PASSWORD'),
        string(credentialsId: 'integration-redis-test', variable: 'REDIS_PASSWORD')
      ]) {
        sh """NETWORK_NAME="${networkName}" docker-compose -p ${env.BUILD_TAG} -f docker-compose-integration.yaml down"""
      }
    }
  }

}
