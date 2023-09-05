import { PartialType } from '@nestjs/swagger';

import { Create{{cookiecutter.className}}Dto } from './create-{{cookiecutter.fileName}}.dto';

export class Update{{cookiecutter.className}}Dto extends PartialType(Create{{cookiecutter.className}}Dto) {}
