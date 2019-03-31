pipeline {
  agent {
    node {
      label 'docker'
    }

  }
  stages {
    stage('stage1') {
      steps {
        git(url: 'https://github.com/danscarf/golang-webapi', branch: 'master')
      }
    }
  }
}