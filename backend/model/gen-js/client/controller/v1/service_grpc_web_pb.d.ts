import * as grpcWeb from 'grpc-web';

import * as client_controller_v1_service_pb from '../../../client/controller/v1/service_pb'; // proto import: "client/controller/v1/service.proto"


export class ControllerServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getListEndpoint(
    request: client_controller_v1_service_pb.GetListEndpointRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: client_controller_v1_service_pb.GetListEndpointResponse) => void
  ): grpcWeb.ClientReadableStream<client_controller_v1_service_pb.GetListEndpointResponse>;

}

export class ControllerServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getListEndpoint(
    request: client_controller_v1_service_pb.GetListEndpointRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<client_controller_v1_service_pb.GetListEndpointResponse>;

}

