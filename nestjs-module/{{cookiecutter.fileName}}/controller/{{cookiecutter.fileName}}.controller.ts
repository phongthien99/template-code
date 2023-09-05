import {
  Body,
  Controller,
  Delete,
  Get,
  Param,
  ParseUUIDPipe,
  Patch,
  Post,
  Query,
  UseInterceptors,
  UsePipes,
  ValidationPipe,
} from '@nestjs/common';
import { ApiBasicAuth, ApiOperation, ApiTags } from '@nestjs/swagger';
import { NotFoundI18nInterceptor } from '@sigmaott/common';
import { AppId } from '@sigmaott/core';
import { CollectionDto, CollectionValidationPipe } from '@sigmaott/paginate';
import { I18n, I18nContext } from 'nestjs-i18n';
import { SanitizeMongooseModelInterceptor } from 'nestjs-mongoose-exclude';

import { Create{{cookiecutter.className}}Dto } from '../dto/create-{{cookiecutter.fileName}}.dto';
import { Update{{cookiecutter.className}}Dto } from '../dto/update-{{cookiecutter.fileName}}.dto';
import { {{cookiecutter.className}}, {{cookiecutter.className}}Collection } from '../entities/{{cookiecutter.fileName}}.entity';
import { {{cookiecutter.className}}Service } from '../service/{{cookiecutter.fileName}}.service';

@Controller('{{cookiecutter.fileName}}')
@ApiTags('{{cookiecutter.className}}')
@ApiBasicAuth()
@UseInterceptors(
  new SanitizeMongooseModelInterceptor({
    excludeMongooseId: true,
    excludeMongooseV: true,
  })
)
export class {{cookiecutter.className}}Controller {
  constructor(private readonly {{cookiecutter.fieldPrefix}}Service: {{cookiecutter.className}}Service) {}

  @Post('/')
  @ApiOperation({
    summary: 'Create {{cookiecutter.className}}',
    description: 'Create {{cookiecutter.className}}',
  })
  @UsePipes(new ValidationPipe({ transform: true, whitelist: true }))
  async create(
    @AppId() appId: string,
    @Body() payload: Create{{cookiecutter.className}}Dto,
    @I18n() i18n: I18nContext
  ): Promise<{{cookiecutter.className}}> {
    return this.{{cookiecutter.fieldPrefix}}Service.create(appId, payload, i18n);
  }

  @Get('/')
  @ApiOperation({
    summary: 'Get {{cookiecutter.className}}s',
    description: 'Get {{cookiecutter.className}}s',
  })
  async list(
    @AppId() appId: string,
    @Query(new CollectionValidationPipe({{cookiecutter.className}}))
    collectionDto: CollectionDto
  ): Promise<{{cookiecutter.className}}Collection> {
    return this.{{cookiecutter.fieldPrefix}}Service.findAll(appId, collectionDto);
  }

  @Get(':id')
  @ApiOperation({
    summary: 'Get {{cookiecutter.className}}  by Id',
    description: 'Get {{cookiecutter.className}}  by Id',
  })
  @UseInterceptors(new NotFoundI18nInterceptor('error.notFound'))
  async getOne(
    @AppId() appId: string,
    @Param('id', new ParseUUIDPipe()) id: string,
    @I18n() i18n: I18nContext
  ): Promise<{{cookiecutter.className}}> {
    const res = await this.{{cookiecutter.fieldPrefix}}Service.findOne(id, appId, i18n);

    return res;
  }

  @Patch(':id')
  @ApiOperation({
    summary: 'Update {{cookiecutter.className}}  by Id',
    description: 'Update {{cookiecutter.className}}  by Id',
  })
  @UsePipes(new ValidationPipe({ transform: true, whitelist: true }))
  @UseInterceptors(new NotFoundI18nInterceptor('error.notFound'))
  updateById(
    @AppId() appId: string,
    @Param('id', new ParseUUIDPipe()) id: string,
    @Body() payload: Update{{cookiecutter.className}}Dto,
    @I18n() i18n: I18nContext
  ) {
    return this.{{cookiecutter.fieldPrefix}}Service.updateOne(
      appId,
      id,
      payload,
      i18n
    );
  }

  @Delete(':id')
  @ApiOperation({
    summary: 'Delete {{cookiecutter.className}} by id',
    description: 'Delete {{cookiecutter.className}} by id',
  })
  @UseInterceptors(new NotFoundI18nInterceptor('error.notFound'))
  async deleteOne(
    @AppId() appId: string,
    @Param('id', new ParseUUIDPipe()) id: string,
    @I18n() i18n: I18nContext
  ): Promise<{{cookiecutter.className}}> {
    const res = await this.{{cookiecutter.fieldPrefix}}Service.deleteOne(appId, id, i18n);

    return res;
  }

}
