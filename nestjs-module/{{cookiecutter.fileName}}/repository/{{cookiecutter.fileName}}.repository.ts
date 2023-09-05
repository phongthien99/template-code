import { InjectModel } from '@nestjs/mongoose';
import { MongooseRepositoryWithPaginate } from '@sigmaott/core';
import { Model } from 'mongoose';

import { {{cookiecutter.className}}, {{cookiecutter.className}}Document } from '../entities/{{cookiecutter.fileName}}.entity';

export class {{cookiecutter.className}}Repository extends MongooseRepositoryWithPaginate<{{cookiecutter.className}}> {
  constructor(
    @InjectModel({{cookiecutter.className}}.name)
    private readonly repository: Model<{{cookiecutter.className}}Document>
  ) {
    super(repository);
  }
}
