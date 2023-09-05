import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';

import { {{cookiecutter.className}}Controller } from './controller/{{cookiecutter.fileName}}.controller';
import { {{cookiecutter.className}}, {{cookiecutter.className}}Schema } from './entities/{{cookiecutter.fileName}}.entity';
import { {{cookiecutter.className}}Repository } from './repository/{{cookiecutter.fileName}}.repository';
import { {{cookiecutter.className}}Service } from './service/{{cookiecutter.fileName}}.service';

@Module({
  imports: [
    MongooseModule.forFeature([
      { name: {{cookiecutter.fileName}}.name, schema: {{cookiecutter.fileName}}Schema },
    ]),
  ],
  controllers: [{{cookiecutter.fileName}}Controller],
  providers: [{{cookiecutter.fileName}}Service, {{cookiecutter.fileName}}Repository],
  exports: [],
})
export class TemplateModule {}
