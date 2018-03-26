'use strict';
var yeoman = require('yeoman-generator');
var _ = require('lodash');
var chalk = require('chalk');
var yosay = require('yosay');
var toml = require('toml');
var ejs = require('ejs');
var path = require('path');

module.exports = class extends yeoman {
  _getRepoUrl() {
    var destinationPath = process.env.LOCAL_PATH || this.destinationRoot();
    var repoUrl = '';
    var src = path.sep + 'src' + path.sep;
    var index = destinationPath.indexOf(src);
    if (index !== -1) {
      repoUrl = destinationPath.substring(index + src.length);
    }
    return repoUrl;
  }

  initializing() {
    this.props = {};
    this.pathToTemplates = '../../templates';
    this.repoUrl = this._getRepoUrl();
    this.IsContinue = this.repoUrl.length > 0;
  }

  configuring() {
    this.projectName = _.kebabCase(this.appname);
  }

  prompting() {
    // Welcome message
    this.log(
      yosay(
        'Welcome to ' +
          chalk`{bold.rgb(239, 115, 16) YoGo} ` +
          chalk`{bold.rgb(105, 215, 226) GoLang} generator!`
      )
    );

    if (!this.IsContinue) {
      this.log(
        `YoGo generator will only generate project in '${chalk.red(
          'GOPATH/src/<YOUR_PROJECT>'
        )} directory'. Otherwise the generation will be stopped.`
      );
      return;
    }

    var prompts = [
      {
        type: 'list',
        name: 'project',
        message: 'What type of application do you want to create?',
        choices: [
          {
            name: 'Empty Console Application (with "Hello world")',
            value: 'console'
          },
          {
            name: 'REST API microservice',
            value: 'restapi'
          },
          {
            name: 'GO-KIT microservice',
            value: 'gokitapi'
          }
        ]
      },
      {
        type: 'list',
        name: 'dependencytool',
        message: 'What type of dependency management tool do you want to use?',
        choices: [
          {
            name: 'golang/dep - (GO official experiment)',
            value: 'dep'
          },
          {
            name: 'glide',
            value: 'glide'
          }
        ],
        when: props => props.project.indexOf('console') === -1
      },
      {
        type: 'input',
        name: 'packagename',
        message: 'Please type your package name',
        default: 'test',
        when: props =>
          props.project.indexOf('gokitapi') !== -1 ||
          props.project.indexOf('restapi') !== -1
      },
      {
        type: 'checkbox',
        name: 'includeFiles',
        message: 'Which additional files would you like to include?',
        choices: [
          {
            name: '.gitignore',
            value: 'gitignore',
            checked: true
          },
          {
            name: 'README.md',
            value: 'readme',
            checked: true
          },
          {
            name: 'Dockerfile',
            value: 'dockerfile',
            checked: true
          },
          {
            name: 'docker-compose',
            value: 'dockercompose',
            checked: false
          },
          {
            name: 'Makefile',
            value: 'makefile',
            checked: true
          }
        ]
      }
    ];

    return this.prompt(prompts).then(props => {
      this.projectType = props.project;
      this.topicName = props.topicname;
      this.packageName = props.packagename;

      this.dependencyManagementTool =
        props.dependencytool === undefined || props.dependencytool === null
          ? ''
          : props.dependencytool;

      this.includeGitIgnore = _.includes(props.includeFiles, 'gitignore');
      this.includeDockerfile = _.includes(props.includeFiles, 'dockerfile');
      this.includeDockerCompose = _.includes(props.includeFiles, 'dockercompose');
      this.includeMakefile = _.includes(props.includeFiles, 'makefile');
      this.includeReadmeFile = _.includes(props.includeFiles, 'readme');
    });
  }

  writing() {
    if (!this.IsContinue) return;

    switch (this.projectType) {
      case 'console':
        this.fs.copy(
          this.templatePath('console/_main.go'),
          this.destinationPath('main.go')
        );
        break;
      case 'restapi':
        this.fs.copyTpl(
          this.templatePath('restapi/_main.go'),
          this.destinationPath('main.go'),
          {
            repourl: this.repoUrl,
            projectname: this.projectName,
            packagename: this.packageName
          }
        );

        // Copy package name
        this.fs.copyTpl(
          this.templatePath(this.pathToTemplates + '/pkg-rest-endpoint/_handler.go'),
          this.destinationPath('./' + this.packageName + '/endpoint.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        this.fs.copyTpl(
          this.templatePath(this.pathToTemplates + '/pkg-rest-endpoint/_interface.go'),
          this.destinationPath('./' + this.packageName + '/interface.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        this.fs.copyTpl(
          this.templatePath(this.pathToTemplates + '/pkg-rest-endpoint/_model.go'),
          this.destinationPath('./' + this.packageName + '/model.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        this.fs.copyTpl(
          this.templatePath(this.pathToTemplates + '/pkg-rest-endpoint/_repository.go'),
          this.destinationPath('./' + this.packageName + '/repository.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        this.fs.copyTpl(
          this.templatePath(
            this.pathToTemplates + '/pkg-rest-endpoint/_service_tracing.go'
          ),
          this.destinationPath('./' + this.packageName + '/service_tracing.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        this.fs.copyTpl(
          this.templatePath(this.pathToTemplates + '/pkg-rest-endpoint/_service.go'),
          this.destinationPath('./' + this.packageName + '/service.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        break;

      case 'gokitapi':
        this.fs.copyTpl(
          this.templatePath('gokitapi/_main.go'),
          this.destinationPath('main.go'),
          {
            repourl: this.repoUrl,
            projectname: this.projectName,
            packagename: this.packageName
          }
        );

        // Copy package name
        this.fs.copyTpl(
          this.templatePath(this.pathToTemplates + '/pkg-gokit-endpoint/_endpoint.go'),
          this.destinationPath('./' + this.packageName + '/endpoint.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        this.fs.copyTpl(
          this.templatePath(
            this.pathToTemplates + '/pkg-gokit-endpoint/_http_transport.go'
          ),
          this.destinationPath('./' + this.packageName + '/http_transport.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        this.fs.copyTpl(
          this.templatePath(this.pathToTemplates + '/pkg-gokit-endpoint/_model.go'),
          this.destinationPath('./' + this.packageName + '/model.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        this.fs.copyTpl(
          this.templatePath(this.pathToTemplates + '/pkg-gokit-endpoint/_service.go'),
          this.destinationPath('./' + this.packageName + '/service.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        this.fs.copyTpl(
          this.templatePath(
            this.pathToTemplates + '/pkg-gokit-endpoint/_service_tracing.go'
          ),
          this.destinationPath('./' + this.packageName + '/service_tracing.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );

        // Copy utest package name
        this.fs.copyTpl(
          this.templatePath(
            this.pathToTemplates + '/pkg-gokit-endpoint/_service_tracing_test.go'
          ),
          this.destinationPath('./' + this.packageName + '/service_tracing_test.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        this.fs.copyTpl(
          this.templatePath(
            this.pathToTemplates + '/pkg-gokit-endpoint/_endpoint_test.go'
          ),
          this.destinationPath('./' + this.packageName + '/endpoint_test.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        this.fs.copyTpl(
          this.templatePath(
            this.pathToTemplates + '/pkg-gokit-endpoint/_http_transport_test.go'
          ),
          this.destinationPath('./' + this.packageName + '/http_transport_test.go'),
          { projectname: this.projectName, packagename: this.packageName }
        );
        break;
      default:
        this.log('nothing to do');
        break;
    }

    // Copy additional files
    this._copyAdditionalFiles();

    // Copy file of dependency tool
    this._copyFileDependencyTool();

    // Copy config package
    if (this.projectType !== 'console') {
      this._copyConfigPackage();
    }
  }

  _copyFileDependencyTool() {
    if (this.projectType !== 'console') {
      switch (this.dependencyManagementTool) {
        case 'glide':
          this.fs.copyTpl(
            this.templatePath(this.projectType + '/_glide.yaml'),
            this.destinationPath('glide.yaml'),
            {
              projectname: this.projectName,
              packagename: this.packageName,
              repourl: this.repoUrl
            }
          );
          break;
        case 'dep':
          this.fs.copyTpl(
            this.templatePath(this.projectType + '/_Gopkg.toml'),
            this.destinationPath('Gopkg.toml'),
            { projectname: this.projectName, packagename: this.packageName }
          );
          break;
        default:
          this.log('nothing to do');
          break;
      }
    }
  }

  _copyConfigPackage() {
    this.ConfigEnvVariables = [];
    var configTomlPath = this.destinationPath() + '/config.toml';

    // Read config.toml file if it's exist
    if (this.fs.exists(configTomlPath)) {
      try {
        this.ConfigEnvVariables = toml.parse(this.fs.read(configTomlPath));
      } catch (e) {
        this.log('\n');
        this.log(
          chalk.default.red(
            'ERROR : Your config.toml file is not valid. YoGo will create default config package.'
          )
        );
        this.log('\n');
      }
    }

    var arrEnvKeys = Object.keys(this.ConfigEnvVariables);
    if (arrEnvKeys.length > 0) {
      // Copy from destination folder to ./config pkg
      this.fs.copy(
        this.destinationPath('./config.toml'),
        this.destinationPath('./config/config.toml')
      );
    } else {
      this.fs.copyTpl(
        this.templatePath(this.pathToTemplates + '/config/_config.toml'),
        this.destinationPath('./config/config.toml')
      );
    }

    this.fs.copyTpl(
      this.templatePath(this.pathToTemplates + '/config/_config_dynamic.toml.example'),
      this.destinationPath('./config/config.toml.example'),
      { envs: arrEnvKeys }
    );
    this.fs.copyTpl(
      this.templatePath(this.pathToTemplates + '/config/_config_dynamic.go'),
      this.destinationPath('./config/config.go'),
      { _: _, envs: arrEnvKeys }
    );
    this.fs.copyTpl(
      this.templatePath(this.pathToTemplates + '/config/_config_test.go'),
      this.destinationPath('./config/config_test.go')
    );
  }

  _copyAdditionalFiles() {
    // Copy gitignore file
    if (this.includeGitIgnore) {
      this.fs.copyTpl(
        this.templatePath('_gitignore'),
        this.destinationPath('.gitignore'),
        { projectname: this.projectName }
      );
    }

    // Copy Docker file
    if (this.includeDockerfile) {
      this.fs.copyTpl(
        this.templatePath('_Dockerfile'),
        this.destinationPath('Dockerfile'),
        { projectname: this.projectName }
      );
    }

    // Copy docker-compose file
    if (this.includeDockerCompose) {
      this.fs.copyTpl(
        this.templatePath('_docker-compose.yml'),
        this.destinationPath('docker-compose.yml'),
        { projectname: this.projectName }
      );
    }

    // Copy Makefile file
    if (this.includeMakefile) {
      this.fs.copyTpl(this.templatePath('_Makefile'), this.destinationPath('Makefile'), {
        projectname: this.projectName,
        dockerrun: this.docker_run,
        dependencytool: this.dependencyManagementTool
      });
    }

    // Copy README file
    if (this.includeReadmeFile) {
      this.fs.copyTpl(
        this.templatePath('_readme.md'),
        this.destinationPath('README.md'),
        { projectname: this.projectName }
      );
    }
  }

  end() {
    if (!this.IsContinue) return;

    this.log('\n');
    this.log(
      '**********************************************************************************'
    );
    this.log(
      '* Your project is now created, follow the instructions below                     *'
    );
    this.log(
      '**********************************************************************************'
    );

    // Build end message content
    var filePath = this.templatePath('./' + this.projectType + '/_end_log.txt');
    var content = this.fs.read(filePath);
    var result = ejs.compile(content)({
      projectname: this.projectName,
      packagename: this.packageName,
      chalk: chalk,
      dependencytool: this.dependencyManagementTool
    });
    this.log(result);
  }
};
