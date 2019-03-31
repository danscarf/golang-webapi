pipeline {
  agent {
    node {
      label 'docker'
    }

  }
  stages {
    stage('stage1') {
      steps {
        checkout(
          [$class: 'GitSCM',
          branches: [[name: '*/master']],
          doGenerateSubmoduleConfigurations: false,
          extensions: [],
          submoduleCfg: [],
          userRemoteConfigs:
            [
              [
                credentialsId: 'github',
                url: 'https://github.com/danscarf/golang-webapi'
              ]
            ]
          ]
        )
      }
    }
  }
}

