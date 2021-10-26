pipeline {
    agent any
    triggers {
        pollSCM('*/5 * * * *')
    }
    stages {
        stage("Building Images") {
            steps {
                sh "docker-compose up --detach --build"
                sh "docker image tag rblog:v1.5.3 ronaldoafonso/rblog:v1.5.3"
                sh "docker image tag rblog_nginx:v0.0.1 ronaldoafonso/rblog_nginx:v0.0.1"
                sh "docker image push ronaldoafonso/rblog:v1.5.3"
                sh "docker image push ronaldoafonso/rblog_nginx:v0.0.1"
            }
        }
        stage("Deploy Images") {
            steps {
                sh "docker -H tcp://lion.vpn.ronaldoafonso.com.br:2375 stack deploy -c docker-stack.yml rblog"
            }
        }
    }
    post {
        always {
            sh "docker-compose down"
            mail body: "Build: ${env.BUILD_URL}", subject: 'rblog Build Status', to: 'ronaldo@vpn.ronaldoafonso.com.br'
        }
    }
}
