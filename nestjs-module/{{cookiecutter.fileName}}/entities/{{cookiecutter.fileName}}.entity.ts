import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { ApiPropertyOptional, CollectionResponse } from '@sigmaott/core';
import { ExposeQuery } from '@sigmaott/paginate';
import { Document } from 'mongoose';
import * as uuid from 'uuid';

export type {{cookiecutter.className}}Document = {{cookiecutter.className}} & Document;

@Schema({
  timestamps: true,
  toJSON: {
    versionKey: false,
    virtuals: true,
    transform: function (doc, ret) {
      delete ret._id;
      delete ret.__v;
    },
  },
})
export class {{cookiecutter.className}} {

  @Prop({ default: uuid.v4 })
  id?: string;

  @Prop()
  @ExposeQuery({ filterable: true, description: 'created at', sortable: true })
  createdAt?: Date;

  @Prop()
  @ExposeQuery({ filterable: true, description: 'updated at', sortable: true })
  updatedAt?: Date;

  @Prop()
  @ExposeQuery({ filterable: true, description: 'app Id', sortable: true })
  appId?: string;
  
}

export const {{cookiecutter.className}}chema = SchemaFactory.createForClass({{cookiecutter.className}});

export class {{cookiecutter.className}}Collection extends CollectionResponse {
  @ApiPropertyOptional({ type: () => [{{cookiecutter.className}}] })
  data: {{cookiecutter.className}}[];
}

