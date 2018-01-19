'use strict';

var yeoman = require('yeoman-generator');
var _ = require('lodash');
var chalk = require('chalk');
var yosay = require('yosay');
var toml = require('toml');
var ejs = require('ejs');

module.exports = class extends yeoman {
  initializing() {
    this.props = {};
    this.pathToTemplates = '../../templates';
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

    var prompts = [
      {
        type: 'list',
        name: 'packagetype',
        message: 'What type of package would you like to create?',
        choices: [
          {
            name: 'Config package',
            value: 'config'
          },
          {
            name: 'REST API package',
            value: 'pkgrestapi'
          }
        ]
      },
      {
        type: 'input',
        name: 'packagename',
        message: 'Please type your package name',
        default: 'test',
        when: props =>
          props.packagetype.indexOf('pkgrestapi') !== -1 ||
          props.packagetype.indexOf('pkggokitapi') !== -1
      }
    ];

    return this.prompt(prompts).then(props => {
      this.packageType = props.packagetype;
      if (this.packageType === 'config') {
        this.packageName = 'config';
      } else {
        this.packageName = props.packagename;
      }
    });
  }

  writing() {
    switch (this.packageType) {
      case 'config':
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
          this.templatePath(
            this.pathToTemplates + '/config/_config_dynamic.toml.example'
          ),
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

        break;
      case 'pkgrestapi':
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
      default:
        this.log('nothing to do');
        break;
    }
  }

  install() {
    // This.log("install");
  }

  end() {
    this.log('\n');
    this.log(
      '**********************************************************************************'
    );
    this.log(
      '* Your package is now created, follow the instructions below                     *'
    );
    this.log(
      '**********************************************************************************'
    );

    var filePath = '';
    var content = '';
    var result = '';

    switch (this.packageType) {
      case 'config':
        break;
      case 'pkgrestapi':
        // Build content
        filePath = this.templatePath(
          this.pathToTemplates + '/pkg-rest-endpoint/_end_log.txt'
        );
        content = this.fs.read(filePath);
        result = ejs.compile(content)({
          projectname: this.projectName,
          packagename: this.packageName,
          chalk: chalk
        });
        this.log(result);
        break;
      default:
        this.log('nothing to do');
        break;
    }
  }
};
