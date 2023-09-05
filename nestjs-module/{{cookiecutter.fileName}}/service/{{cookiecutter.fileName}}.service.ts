import { BadRequestException, Injectable } from '@nestjs/common';
import { CollectionDto } from '@sigmaott/paginate';
import { I18nContext } from 'nestjs-i18n';

import { Create{{cookiecutter.className}}Dto } from '../dto/create-{{cookiecutter.fileName}}.dto';
import { Update{{cookiecutter.className}}Dto } from '../dto/update-{{cookiecutter.fileName}}.dto';
import { {{cookiecutter.className}}, {{cookiecutter.className}}Collection } from '../entities/{{cookiecutter.fileName}}.entity';
import { {{cookiecutter.className}}Repository } from '../repository/{{cookiecutter.fileName}}.repository';

@Injectable()
export class {{cookiecutter.className}}Service {
  constructor(private readonly {{cookiecutter.fieldPrefix}}Repository: {{cookiecutter.className}}Repository) { }

  async create(
    appId: string,
    payload: Create{{cookiecutter.className}}Dto,
    i18n: I18nContext
  ) {
    
    const modelCreate = {
      ...payload,
      appId:appId
    } as {{cookiecutter.className}};

   
    const res = await this.{{cookiecutter.fieldPrefix}}Repository.create(modelCreate);

    return res;
  }

  async findAll(
    appId: string,
    collectionDto: CollectionDto
  ): Promise< {{cookiecutter.className}}Collection> {
    collectionDto.filter.appId = appId

    return this.{{cookiecutter.fieldPrefix}}Repository.paginate(collectionDto);
  }

  async findOne(id: string, appId: string, i18n: I18nContext) {
   

    const res = await this.{{cookiecutter.fieldPrefix}}Repository.findOne({ id: id});
    if (!res||res.appId !== appId){
      return 
    }
    return res
  }


  async updateOne(
    id: string,
    appId: string,
    payload: Update{{cookiecutter.className}}Dto,
    i18n: I18nContext
  ) {
 

    return await this.{{cookiecutter.fieldPrefix}}Repository.updateOne(
      { id: id, appId: appId },
      payload
    );
  }

  async deleteOne(appId: string, id: string, i18n: I18nContext) {

    const res = await this.{{cookiecutter.fieldPrefix}}Repository.findOne({ id: id });
    if (!res||res.appId !== appId){
      return 
    }

    return await this.{{cookiecutter.fieldPrefix}}Repository.deleteOne({ id: id });


  }

  




}
