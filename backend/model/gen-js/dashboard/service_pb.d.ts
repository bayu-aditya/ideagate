import * as jspb from 'google-protobuf'

import * as core_application_application_pb from '../core/application/application_pb'; // proto import: "core/application/application.proto"
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb'; // proto import: "google/protobuf/struct.proto"


export class GetListApplicationRequest extends jspb.Message {
  getProjectId(): string;
  setProjectId(value: string): GetListApplicationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetListApplicationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetListApplicationRequest): GetListApplicationRequest.AsObject;
  static serializeBinaryToWriter(message: GetListApplicationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetListApplicationRequest;
  static deserializeBinaryFromReader(message: GetListApplicationRequest, reader: jspb.BinaryReader): GetListApplicationRequest;
}

export namespace GetListApplicationRequest {
  export type AsObject = {
    projectId: string,
  }
}

export class GetListApplicationResponse extends jspb.Message {
  getApplicationsList(): Array<core_application_application_pb.Application>;
  setApplicationsList(value: Array<core_application_application_pb.Application>): GetListApplicationResponse;
  clearApplicationsList(): GetListApplicationResponse;
  addApplications(value?: core_application_application_pb.Application, index?: number): core_application_application_pb.Application;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetListApplicationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetListApplicationResponse): GetListApplicationResponse.AsObject;
  static serializeBinaryToWriter(message: GetListApplicationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetListApplicationResponse;
  static deserializeBinaryFromReader(message: GetListApplicationResponse, reader: jspb.BinaryReader): GetListApplicationResponse;
}

export namespace GetListApplicationResponse {
  export type AsObject = {
    applicationsList: Array<core_application_application_pb.Application.AsObject>,
  }
}

export class UpsertApplicationRequest extends jspb.Message {
  getProjectId(): string;
  setProjectId(value: string): UpsertApplicationRequest;

  getApplicationId(): string;
  setApplicationId(value: string): UpsertApplicationRequest;

  getValues(): google_protobuf_struct_pb.Struct | undefined;
  setValues(value?: google_protobuf_struct_pb.Struct): UpsertApplicationRequest;
  hasValues(): boolean;
  clearValues(): UpsertApplicationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpsertApplicationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpsertApplicationRequest): UpsertApplicationRequest.AsObject;
  static serializeBinaryToWriter(message: UpsertApplicationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpsertApplicationRequest;
  static deserializeBinaryFromReader(message: UpsertApplicationRequest, reader: jspb.BinaryReader): UpsertApplicationRequest;
}

export namespace UpsertApplicationRequest {
  export type AsObject = {
    projectId: string,
    applicationId: string,
    values?: google_protobuf_struct_pb.Struct.AsObject,
  }
}

export class UpsertApplicationResponse extends jspb.Message {
  getApplication(): core_application_application_pb.Application | undefined;
  setApplication(value?: core_application_application_pb.Application): UpsertApplicationResponse;
  hasApplication(): boolean;
  clearApplication(): UpsertApplicationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpsertApplicationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpsertApplicationResponse): UpsertApplicationResponse.AsObject;
  static serializeBinaryToWriter(message: UpsertApplicationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpsertApplicationResponse;
  static deserializeBinaryFromReader(message: UpsertApplicationResponse, reader: jspb.BinaryReader): UpsertApplicationResponse;
}

export namespace UpsertApplicationResponse {
  export type AsObject = {
    application?: core_application_application_pb.Application.AsObject,
  }
}

export class DeleteApplicationRequest extends jspb.Message {
  getProjectId(): string;
  setProjectId(value: string): DeleteApplicationRequest;

  getApplicationId(): string;
  setApplicationId(value: string): DeleteApplicationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteApplicationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteApplicationRequest): DeleteApplicationRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteApplicationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteApplicationRequest;
  static deserializeBinaryFromReader(message: DeleteApplicationRequest, reader: jspb.BinaryReader): DeleteApplicationRequest;
}

export namespace DeleteApplicationRequest {
  export type AsObject = {
    projectId: string,
    applicationId: string,
  }
}

export class DeleteApplicationResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteApplicationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteApplicationResponse): DeleteApplicationResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteApplicationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteApplicationResponse;
  static deserializeBinaryFromReader(message: DeleteApplicationResponse, reader: jspb.BinaryReader): DeleteApplicationResponse;
}

export namespace DeleteApplicationResponse {
  export type AsObject = {
  }
}

