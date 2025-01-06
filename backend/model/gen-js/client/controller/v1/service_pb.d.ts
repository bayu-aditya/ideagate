import * as jspb from 'google-protobuf'

import * as core_endpoint_endpoint_pb from '../../../core/endpoint/endpoint_pb'; // proto import: "core/endpoint/endpoint.proto"


export class GetListEndpointRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetListEndpointRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetListEndpointRequest): GetListEndpointRequest.AsObject;
  static serializeBinaryToWriter(message: GetListEndpointRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetListEndpointRequest;
  static deserializeBinaryFromReader(message: GetListEndpointRequest, reader: jspb.BinaryReader): GetListEndpointRequest;
}

export namespace GetListEndpointRequest {
  export type AsObject = {
  }
}

export class GetListEndpointResponse extends jspb.Message {
  getEndpointsList(): Array<core_endpoint_endpoint_pb.Endpoint>;
  setEndpointsList(value: Array<core_endpoint_endpoint_pb.Endpoint>): GetListEndpointResponse;
  clearEndpointsList(): GetListEndpointResponse;
  addEndpoints(value?: core_endpoint_endpoint_pb.Endpoint, index?: number): core_endpoint_endpoint_pb.Endpoint;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetListEndpointResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetListEndpointResponse): GetListEndpointResponse.AsObject;
  static serializeBinaryToWriter(message: GetListEndpointResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetListEndpointResponse;
  static deserializeBinaryFromReader(message: GetListEndpointResponse, reader: jspb.BinaryReader): GetListEndpointResponse;
}

export namespace GetListEndpointResponse {
  export type AsObject = {
    endpointsList: Array<core_endpoint_endpoint_pb.Endpoint.AsObject>,
  }
}

