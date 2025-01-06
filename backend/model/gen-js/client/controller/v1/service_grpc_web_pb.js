/**
 * @fileoverview gRPC-Web generated client stub for v1
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v0.0.0
// source: client/controller/v1/service.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var core_endpoint_endpoint_pb = require('../../../core/endpoint/endpoint_pb.js')
const proto = {};
proto.v1 = require('./service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.v1.ControllerServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.v1.ControllerServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.v1.GetListEndpointRequest,
 *   !proto.v1.GetListEndpointResponse>}
 */
const methodDescriptor_ControllerService_GetListEndpoint = new grpc.web.MethodDescriptor(
  '/v1.ControllerService/GetListEndpoint',
  grpc.web.MethodType.UNARY,
  proto.v1.GetListEndpointRequest,
  proto.v1.GetListEndpointResponse,
  /**
   * @param {!proto.v1.GetListEndpointRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.v1.GetListEndpointResponse.deserializeBinary
);


/**
 * @param {!proto.v1.GetListEndpointRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.v1.GetListEndpointResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.v1.GetListEndpointResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.v1.ControllerServiceClient.prototype.getListEndpoint =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/v1.ControllerService/GetListEndpoint',
      request,
      metadata || {},
      methodDescriptor_ControllerService_GetListEndpoint,
      callback);
};


/**
 * @param {!proto.v1.GetListEndpointRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.v1.GetListEndpointResponse>}
 *     Promise that resolves to the response
 */
proto.v1.ControllerServicePromiseClient.prototype.getListEndpoint =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/v1.ControllerService/GetListEndpoint',
      request,
      metadata || {},
      methodDescriptor_ControllerService_GetListEndpoint);
};


module.exports = proto.v1;

