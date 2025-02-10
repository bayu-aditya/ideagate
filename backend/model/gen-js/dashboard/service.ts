// @generated by protobuf-ts 2.9.4
// @generated from protobuf file "dashboard/service.proto" (package "dashboard", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Struct } from "../google/protobuf/struct";
import { Application } from "../core/application/application";
/**
 * @generated from protobuf message dashboard.GetListApplicationRequest
 */
export interface GetListApplicationRequest {
    /**
     * @generated from protobuf field: string project_id = 1;
     */
    projectId: string;
}
/**
 * @generated from protobuf message dashboard.GetListApplicationResponse
 */
export interface GetListApplicationResponse {
    /**
     * @generated from protobuf field: repeated application.Application applications = 1;
     */
    applications: Application[];
}
/**
 * @generated from protobuf message dashboard.UpsertApplicationRequest
 */
export interface UpsertApplicationRequest {
    /**
     * @generated from protobuf field: string project_id = 1;
     */
    projectId: string;
    /**
     * @generated from protobuf field: string application_id = 2;
     */
    applicationId: string;
    /**
     * @generated from protobuf field: google.protobuf.Struct values = 3;
     */
    values?: Struct;
}
/**
 * @generated from protobuf message dashboard.UpsertApplicationResponse
 */
export interface UpsertApplicationResponse {
    /**
     * @generated from protobuf field: application.Application application = 1;
     */
    application?: Application;
}
/**
 * @generated from protobuf message dashboard.DeleteApplicationRequest
 */
export interface DeleteApplicationRequest {
    /**
     * @generated from protobuf field: string project_id = 1;
     */
    projectId: string;
    /**
     * @generated from protobuf field: string application_id = 2;
     */
    applicationId: string;
}
/**
 * @generated from protobuf message dashboard.DeleteApplicationResponse
 */
export interface DeleteApplicationResponse {
}
// @generated message type with reflection information, may provide speed optimized methods
class GetListApplicationRequest$Type extends MessageType<GetListApplicationRequest> {
    constructor() {
        super("dashboard.GetListApplicationRequest", [
            { no: 1, name: "project_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<GetListApplicationRequest>): GetListApplicationRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.projectId = "";
        if (value !== undefined)
            reflectionMergePartial<GetListApplicationRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetListApplicationRequest): GetListApplicationRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string project_id */ 1:
                    message.projectId = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetListApplicationRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string project_id = 1; */
        if (message.projectId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.projectId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.GetListApplicationRequest
 */
export const GetListApplicationRequest = new GetListApplicationRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetListApplicationResponse$Type extends MessageType<GetListApplicationResponse> {
    constructor() {
        super("dashboard.GetListApplicationResponse", [
            { no: 1, name: "applications", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Application }
        ]);
    }
    create(value?: PartialMessage<GetListApplicationResponse>): GetListApplicationResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.applications = [];
        if (value !== undefined)
            reflectionMergePartial<GetListApplicationResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetListApplicationResponse): GetListApplicationResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated application.Application applications */ 1:
                    message.applications.push(Application.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetListApplicationResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated application.Application applications = 1; */
        for (let i = 0; i < message.applications.length; i++)
            Application.internalBinaryWrite(message.applications[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.GetListApplicationResponse
 */
export const GetListApplicationResponse = new GetListApplicationResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpsertApplicationRequest$Type extends MessageType<UpsertApplicationRequest> {
    constructor() {
        super("dashboard.UpsertApplicationRequest", [
            { no: 1, name: "project_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "application_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "values", kind: "message", T: () => Struct }
        ]);
    }
    create(value?: PartialMessage<UpsertApplicationRequest>): UpsertApplicationRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.projectId = "";
        message.applicationId = "";
        if (value !== undefined)
            reflectionMergePartial<UpsertApplicationRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UpsertApplicationRequest): UpsertApplicationRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string project_id */ 1:
                    message.projectId = reader.string();
                    break;
                case /* string application_id */ 2:
                    message.applicationId = reader.string();
                    break;
                case /* google.protobuf.Struct values */ 3:
                    message.values = Struct.internalBinaryRead(reader, reader.uint32(), options, message.values);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: UpsertApplicationRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string project_id = 1; */
        if (message.projectId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.projectId);
        /* string application_id = 2; */
        if (message.applicationId !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.applicationId);
        /* google.protobuf.Struct values = 3; */
        if (message.values)
            Struct.internalBinaryWrite(message.values, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.UpsertApplicationRequest
 */
export const UpsertApplicationRequest = new UpsertApplicationRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpsertApplicationResponse$Type extends MessageType<UpsertApplicationResponse> {
    constructor() {
        super("dashboard.UpsertApplicationResponse", [
            { no: 1, name: "application", kind: "message", T: () => Application }
        ]);
    }
    create(value?: PartialMessage<UpsertApplicationResponse>): UpsertApplicationResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<UpsertApplicationResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UpsertApplicationResponse): UpsertApplicationResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* application.Application application */ 1:
                    message.application = Application.internalBinaryRead(reader, reader.uint32(), options, message.application);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: UpsertApplicationResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* application.Application application = 1; */
        if (message.application)
            Application.internalBinaryWrite(message.application, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.UpsertApplicationResponse
 */
export const UpsertApplicationResponse = new UpsertApplicationResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteApplicationRequest$Type extends MessageType<DeleteApplicationRequest> {
    constructor() {
        super("dashboard.DeleteApplicationRequest", [
            { no: 1, name: "project_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "application_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<DeleteApplicationRequest>): DeleteApplicationRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.projectId = "";
        message.applicationId = "";
        if (value !== undefined)
            reflectionMergePartial<DeleteApplicationRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteApplicationRequest): DeleteApplicationRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string project_id */ 1:
                    message.projectId = reader.string();
                    break;
                case /* string application_id */ 2:
                    message.applicationId = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: DeleteApplicationRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string project_id = 1; */
        if (message.projectId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.projectId);
        /* string application_id = 2; */
        if (message.applicationId !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.applicationId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.DeleteApplicationRequest
 */
export const DeleteApplicationRequest = new DeleteApplicationRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteApplicationResponse$Type extends MessageType<DeleteApplicationResponse> {
    constructor() {
        super("dashboard.DeleteApplicationResponse", []);
    }
    create(value?: PartialMessage<DeleteApplicationResponse>): DeleteApplicationResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<DeleteApplicationResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteApplicationResponse): DeleteApplicationResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: DeleteApplicationResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dashboard.DeleteApplicationResponse
 */
export const DeleteApplicationResponse = new DeleteApplicationResponse$Type();
/**
 * @generated ServiceType for protobuf service dashboard.DashboardService
 */
export const DashboardService = new ServiceType("dashboard.DashboardService", [
    { name: "GetListApplication", options: {}, I: GetListApplicationRequest, O: GetListApplicationResponse },
    { name: "UpsertApplication", options: {}, I: UpsertApplicationRequest, O: UpsertApplicationResponse },
    { name: "DeleteApplication", options: {}, I: DeleteApplicationRequest, O: DeleteApplicationResponse }
]);
