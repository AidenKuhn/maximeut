// @generated by protobuf-ts 2.0.0-alpha.12 with parameters generate_dependencies,optimize_code_size
// @generated from protobuf file "service-example.proto" (package "spec", syntax proto3)
// tslint:disable
import { RpcTransport } from "@protobuf-ts/runtime-rpc";
import { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { ExampleService } from "./service-example";
import { DuplexStreamingCall } from "@protobuf-ts/runtime-rpc";
import { ClientStreamingCall } from "@protobuf-ts/runtime-rpc";
import { ServerStreamingCall } from "@protobuf-ts/runtime-rpc";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import { ExampleResponse } from "./service-example";
import { ExampleRequest } from "./service-example";
import { UnaryCall } from "@protobuf-ts/runtime-rpc";
import { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service spec.ExampleService
 */
export interface IExampleServiceClient {
    /**
     * @generated from protobuf rpc: Unary(spec.ExampleRequest) returns (spec.ExampleResponse);
     */
    unary(input: ExampleRequest, options?: RpcOptions): UnaryCall<ExampleRequest, ExampleResponse>;
    /**
     * @generated from protobuf rpc: ServerStream(spec.ExampleRequest) returns (stream spec.ExampleResponse);
     */
    serverStream(input: ExampleRequest, options?: RpcOptions): ServerStreamingCall<ExampleRequest, ExampleResponse>;
    /**
     * @generated from protobuf rpc: ClientStream(stream spec.ExampleRequest) returns (spec.ExampleResponse);
     */
    clientStream(options?: RpcOptions): ClientStreamingCall<ExampleRequest, ExampleResponse>;
    /**
     * @generated from protobuf rpc: Bidi(stream spec.ExampleRequest) returns (stream spec.ExampleResponse);
     */
    bidi(options?: RpcOptions): DuplexStreamingCall<ExampleRequest, ExampleResponse>;
}
/**
 * @generated from protobuf service spec.ExampleService
 */
export class ExampleServiceClient implements IExampleServiceClient, ServiceInfo {
    typeName = ExampleService.typeName;
    methods = ExampleService.methods;
    options = ExampleService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: Unary(spec.ExampleRequest) returns (spec.ExampleResponse);
     */
    unary(input: ExampleRequest, options?: RpcOptions): UnaryCall<ExampleRequest, ExampleResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ExampleRequest, ExampleResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ServerStream(spec.ExampleRequest) returns (stream spec.ExampleResponse);
     */
    serverStream(input: ExampleRequest, options?: RpcOptions): ServerStreamingCall<ExampleRequest, ExampleResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<ExampleRequest, ExampleResponse>("serverStreaming", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ClientStream(stream spec.ExampleRequest) returns (spec.ExampleResponse);
     */
    clientStream(options?: RpcOptions): ClientStreamingCall<ExampleRequest, ExampleResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<ExampleRequest, ExampleResponse>("clientStreaming", this._transport, method, opt);
    }
    /**
     * @generated from protobuf rpc: Bidi(stream spec.ExampleRequest) returns (stream spec.ExampleResponse);
     */
    bidi(options?: RpcOptions): DuplexStreamingCall<ExampleRequest, ExampleResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<ExampleRequest, ExampleResponse>("duplex", this._transport, method, opt);
    }
}
