import * as jspb from 'google-protobuf'

import * as client_controller_v1_pod_pb from '../../../client/controller/v1/pod_pb'; // proto import: "client/controller/v1/pod.proto"


export class GetListPodRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetListPodRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetListPodRequest): GetListPodRequest.AsObject;
  static serializeBinaryToWriter(message: GetListPodRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetListPodRequest;
  static deserializeBinaryFromReader(message: GetListPodRequest, reader: jspb.BinaryReader): GetListPodRequest;
}

export namespace GetListPodRequest {
  export type AsObject = {
  }
}

export class GetListPodResponse extends jspb.Message {
  getPodsList(): Array<client_controller_v1_pod_pb.Pod>;
  setPodsList(value: Array<client_controller_v1_pod_pb.Pod>): GetListPodResponse;
  clearPodsList(): GetListPodResponse;
  addPods(value?: client_controller_v1_pod_pb.Pod, index?: number): client_controller_v1_pod_pb.Pod;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetListPodResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetListPodResponse): GetListPodResponse.AsObject;
  static serializeBinaryToWriter(message: GetListPodResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetListPodResponse;
  static deserializeBinaryFromReader(message: GetListPodResponse, reader: jspb.BinaryReader): GetListPodResponse;
}

export namespace GetListPodResponse {
  export type AsObject = {
    podsList: Array<client_controller_v1_pod_pb.Pod.AsObject>,
  }
}

