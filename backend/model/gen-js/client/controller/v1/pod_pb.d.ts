import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"


export class Pod extends jspb.Message {
  getName(): string;
  setName(value: string): Pod;

  getNamespace(): string;
  setNamespace(value: string): Pod;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Pod;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Pod;

  getStatus(): PodStatus;
  setStatus(value: PodStatus): Pod;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Pod.AsObject;
  static toObject(includeInstance: boolean, msg: Pod): Pod.AsObject;
  static serializeBinaryToWriter(message: Pod, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Pod;
  static deserializeBinaryFromReader(message: Pod, reader: jspb.BinaryReader): Pod;
}

export namespace Pod {
  export type AsObject = {
    name: string,
    namespace: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    status: PodStatus,
  }
}

export enum PodStatus { 
  POD_STATUS_UNSPECIFIED = 0,
  POD_STATUS_PENDING = 1,
  POD_STATUS_RUNNING = 2,
  POD_STATUS_SUCCEEDED = 3,
  POD_STATUS_FAILED = 4,
}
