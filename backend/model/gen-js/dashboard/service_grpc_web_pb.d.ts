import * as grpcWeb from 'grpc-web';

import * as dashboard_service_pb from '../dashboard/service_pb'; // proto import: "dashboard/service.proto"


export class DashboardServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getListApplication(
    request: dashboard_service_pb.GetListApplicationRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: dashboard_service_pb.GetListApplicationResponse) => void
  ): grpcWeb.ClientReadableStream<dashboard_service_pb.GetListApplicationResponse>;

  upsertApplication(
    request: dashboard_service_pb.UpsertApplicationRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: dashboard_service_pb.UpsertApplicationResponse) => void
  ): grpcWeb.ClientReadableStream<dashboard_service_pb.UpsertApplicationResponse>;

  deleteApplication(
    request: dashboard_service_pb.DeleteApplicationRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: dashboard_service_pb.DeleteApplicationResponse) => void
  ): grpcWeb.ClientReadableStream<dashboard_service_pb.DeleteApplicationResponse>;

}

export class DashboardServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getListApplication(
    request: dashboard_service_pb.GetListApplicationRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<dashboard_service_pb.GetListApplicationResponse>;

  upsertApplication(
    request: dashboard_service_pb.UpsertApplicationRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<dashboard_service_pb.UpsertApplicationResponse>;

  deleteApplication(
    request: dashboard_service_pb.DeleteApplicationRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<dashboard_service_pb.DeleteApplicationResponse>;

}

