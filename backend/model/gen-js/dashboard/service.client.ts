// @generated by protobuf-ts 2.9.4
// @generated from protobuf file "dashboard/service.proto" (package "dashboard", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { DashboardService } from "./service";
import type { DeleteApplicationResponse } from "./service";
import type { DeleteApplicationRequest } from "./service";
import type { UpdateApplicationResponse } from "./service";
import type { UpdateApplicationRequest } from "./service";
import type { CreateApplicationResponse } from "./service";
import type { CreateApplicationRequest } from "./service";
import type { GetListApplicationResponse } from "./service";
import type { GetListApplicationRequest } from "./service";
import type { DeleteProjectResponse } from "./project";
import type { DeleteProjectRequest } from "./project";
import type { UpdateProjectResponse } from "./project";
import type { UpdateProjectRequest } from "./project";
import type { CreateProjectResponse } from "./project";
import type { CreateProjectRequest } from "./project";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { GetListProjectResponse } from "./project";
import type { GetListProjectRequest } from "./project";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service dashboard.DashboardService
 */
export interface IDashboardServiceClient {
    /**
     * @generated from protobuf rpc: GetListProject(dashboard.GetListProjectRequest) returns (dashboard.GetListProjectResponse);
     */
    getListProject(input: GetListProjectRequest, options?: RpcOptions): UnaryCall<GetListProjectRequest, GetListProjectResponse>;
    /**
     * @generated from protobuf rpc: CreateProject(dashboard.CreateProjectRequest) returns (dashboard.CreateProjectResponse);
     */
    createProject(input: CreateProjectRequest, options?: RpcOptions): UnaryCall<CreateProjectRequest, CreateProjectResponse>;
    /**
     * @generated from protobuf rpc: UpdateProject(dashboard.UpdateProjectRequest) returns (dashboard.UpdateProjectResponse);
     */
    updateProject(input: UpdateProjectRequest, options?: RpcOptions): UnaryCall<UpdateProjectRequest, UpdateProjectResponse>;
    /**
     * @generated from protobuf rpc: DeleteProject(dashboard.DeleteProjectRequest) returns (dashboard.DeleteProjectResponse);
     */
    deleteProject(input: DeleteProjectRequest, options?: RpcOptions): UnaryCall<DeleteProjectRequest, DeleteProjectResponse>;
    /**
     * @generated from protobuf rpc: GetListApplication(dashboard.GetListApplicationRequest) returns (dashboard.GetListApplicationResponse);
     */
    getListApplication(input: GetListApplicationRequest, options?: RpcOptions): UnaryCall<GetListApplicationRequest, GetListApplicationResponse>;
    /**
     * @generated from protobuf rpc: CreateApplication(dashboard.CreateApplicationRequest) returns (dashboard.CreateApplicationResponse);
     */
    createApplication(input: CreateApplicationRequest, options?: RpcOptions): UnaryCall<CreateApplicationRequest, CreateApplicationResponse>;
    /**
     * @generated from protobuf rpc: UpdateApplication(dashboard.UpdateApplicationRequest) returns (dashboard.UpdateApplicationResponse);
     */
    updateApplication(input: UpdateApplicationRequest, options?: RpcOptions): UnaryCall<UpdateApplicationRequest, UpdateApplicationResponse>;
    /**
     * @generated from protobuf rpc: DeleteApplication(dashboard.DeleteApplicationRequest) returns (dashboard.DeleteApplicationResponse);
     */
    deleteApplication(input: DeleteApplicationRequest, options?: RpcOptions): UnaryCall<DeleteApplicationRequest, DeleteApplicationResponse>;
}
/**
 * @generated from protobuf service dashboard.DashboardService
 */
export class DashboardServiceClient implements IDashboardServiceClient, ServiceInfo {
    typeName = DashboardService.typeName;
    methods = DashboardService.methods;
    options = DashboardService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: GetListProject(dashboard.GetListProjectRequest) returns (dashboard.GetListProjectResponse);
     */
    getListProject(input: GetListProjectRequest, options?: RpcOptions): UnaryCall<GetListProjectRequest, GetListProjectResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetListProjectRequest, GetListProjectResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: CreateProject(dashboard.CreateProjectRequest) returns (dashboard.CreateProjectResponse);
     */
    createProject(input: CreateProjectRequest, options?: RpcOptions): UnaryCall<CreateProjectRequest, CreateProjectResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateProjectRequest, CreateProjectResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: UpdateProject(dashboard.UpdateProjectRequest) returns (dashboard.UpdateProjectResponse);
     */
    updateProject(input: UpdateProjectRequest, options?: RpcOptions): UnaryCall<UpdateProjectRequest, UpdateProjectResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateProjectRequest, UpdateProjectResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: DeleteProject(dashboard.DeleteProjectRequest) returns (dashboard.DeleteProjectResponse);
     */
    deleteProject(input: DeleteProjectRequest, options?: RpcOptions): UnaryCall<DeleteProjectRequest, DeleteProjectResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteProjectRequest, DeleteProjectResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetListApplication(dashboard.GetListApplicationRequest) returns (dashboard.GetListApplicationResponse);
     */
    getListApplication(input: GetListApplicationRequest, options?: RpcOptions): UnaryCall<GetListApplicationRequest, GetListApplicationResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetListApplicationRequest, GetListApplicationResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: CreateApplication(dashboard.CreateApplicationRequest) returns (dashboard.CreateApplicationResponse);
     */
    createApplication(input: CreateApplicationRequest, options?: RpcOptions): UnaryCall<CreateApplicationRequest, CreateApplicationResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateApplicationRequest, CreateApplicationResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: UpdateApplication(dashboard.UpdateApplicationRequest) returns (dashboard.UpdateApplicationResponse);
     */
    updateApplication(input: UpdateApplicationRequest, options?: RpcOptions): UnaryCall<UpdateApplicationRequest, UpdateApplicationResponse> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateApplicationRequest, UpdateApplicationResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: DeleteApplication(dashboard.DeleteApplicationRequest) returns (dashboard.DeleteApplicationResponse);
     */
    deleteApplication(input: DeleteApplicationRequest, options?: RpcOptions): UnaryCall<DeleteApplicationRequest, DeleteApplicationResponse> {
        const method = this.methods[7], opt = this._transport.mergeOptions(options);
        return stackIntercept<DeleteApplicationRequest, DeleteApplicationResponse>("unary", this._transport, method, opt, input);
    }
}
