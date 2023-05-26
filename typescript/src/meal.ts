/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "meal";

export enum Time {
  SHORT = 0,
  MEDIUM = 1,
  LONG = 2,
  UNRECOGNIZED = -1,
}

export function timeFromJSON(object: any): Time {
  switch (object) {
    case 0:
    case "SHORT":
      return Time.SHORT;
    case 1:
    case "MEDIUM":
      return Time.MEDIUM;
    case 2:
    case "LONG":
      return Time.LONG;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Time.UNRECOGNIZED;
  }
}

export function timeToJSON(object: Time): string {
  switch (object) {
    case Time.SHORT:
      return "SHORT";
    case Time.MEDIUM:
      return "MEDIUM";
    case Time.LONG:
      return "LONG";
    case Time.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum Status {
  SUCCESS = 0,
  ERROR = 1,
  UNRECOGNIZED = -1,
}

export function statusFromJSON(object: any): Status {
  switch (object) {
    case 0:
    case "SUCCESS":
      return Status.SUCCESS;
    case 1:
    case "ERROR":
      return Status.ERROR;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Status.UNRECOGNIZED;
  }
}

export function statusToJSON(object: Status): string {
  switch (object) {
    case Status.SUCCESS:
      return "SUCCESS";
    case Status.ERROR:
      return "ERROR";
    case Status.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Meal {
  id: string;
  name: string;
  duration: Time;
}

/** create */
export interface CreateMealRequest {
  name: string;
  duration: Time;
}

export interface CreateMealResponse {
  status: Status;
  error: string;
  data: Meal | undefined;
}

/** find one */
export interface FindOneRequest {
  id: string;
}

export interface FindOneResponse {
  status: Status;
  error: string;
  data: Meal | undefined;
}

function createBaseMeal(): Meal {
  return { id: "", name: "", duration: 0 };
}

export const Meal = {
  encode(message: Meal, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.duration !== 0) {
      writer.uint32(24).int32(message.duration);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Meal {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMeal();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.duration = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Meal {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      name: isSet(object.name) ? String(object.name) : "",
      duration: isSet(object.duration) ? timeFromJSON(object.duration) : 0,
    };
  },

  toJSON(message: Meal): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.duration !== undefined && (obj.duration = timeToJSON(message.duration));
    return obj;
  },

  create<I extends Exact<DeepPartial<Meal>, I>>(base?: I): Meal {
    return Meal.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Meal>, I>>(object: I): Meal {
    const message = createBaseMeal();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.duration = object.duration ?? 0;
    return message;
  },
};

function createBaseCreateMealRequest(): CreateMealRequest {
  return { name: "", duration: 0 };
}

export const CreateMealRequest = {
  encode(message: CreateMealRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.duration !== 0) {
      writer.uint32(16).int32(message.duration);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateMealRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateMealRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.duration = reader.int32() as any;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateMealRequest {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      duration: isSet(object.duration) ? timeFromJSON(object.duration) : 0,
    };
  },

  toJSON(message: CreateMealRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.duration !== undefined && (obj.duration = timeToJSON(message.duration));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateMealRequest>, I>>(base?: I): CreateMealRequest {
    return CreateMealRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateMealRequest>, I>>(object: I): CreateMealRequest {
    const message = createBaseCreateMealRequest();
    message.name = object.name ?? "";
    message.duration = object.duration ?? 0;
    return message;
  },
};

function createBaseCreateMealResponse(): CreateMealResponse {
  return { status: 0, error: "", data: undefined };
}

export const CreateMealResponse = {
  encode(message: CreateMealResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    if (message.error !== "") {
      writer.uint32(18).string(message.error);
    }
    if (message.data !== undefined) {
      Meal.encode(message.data, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateMealResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateMealResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.status = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.error = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.data = Meal.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateMealResponse {
    return {
      status: isSet(object.status) ? statusFromJSON(object.status) : 0,
      error: isSet(object.error) ? String(object.error) : "",
      data: isSet(object.data) ? Meal.fromJSON(object.data) : undefined,
    };
  },

  toJSON(message: CreateMealResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = statusToJSON(message.status));
    message.error !== undefined && (obj.error = message.error);
    message.data !== undefined && (obj.data = message.data ? Meal.toJSON(message.data) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateMealResponse>, I>>(base?: I): CreateMealResponse {
    return CreateMealResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateMealResponse>, I>>(object: I): CreateMealResponse {
    const message = createBaseCreateMealResponse();
    message.status = object.status ?? 0;
    message.error = object.error ?? "";
    message.data = (object.data !== undefined && object.data !== null) ? Meal.fromPartial(object.data) : undefined;
    return message;
  },
};

function createBaseFindOneRequest(): FindOneRequest {
  return { id: "" };
}

export const FindOneRequest = {
  encode(message: FindOneRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FindOneRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFindOneRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FindOneRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: FindOneRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  create<I extends Exact<DeepPartial<FindOneRequest>, I>>(base?: I): FindOneRequest {
    return FindOneRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<FindOneRequest>, I>>(object: I): FindOneRequest {
    const message = createBaseFindOneRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseFindOneResponse(): FindOneResponse {
  return { status: 0, error: "", data: undefined };
}

export const FindOneResponse = {
  encode(message: FindOneResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    if (message.error !== "") {
      writer.uint32(18).string(message.error);
    }
    if (message.data !== undefined) {
      Meal.encode(message.data, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FindOneResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFindOneResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.status = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.error = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.data = Meal.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FindOneResponse {
    return {
      status: isSet(object.status) ? statusFromJSON(object.status) : 0,
      error: isSet(object.error) ? String(object.error) : "",
      data: isSet(object.data) ? Meal.fromJSON(object.data) : undefined,
    };
  },

  toJSON(message: FindOneResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = statusToJSON(message.status));
    message.error !== undefined && (obj.error = message.error);
    message.data !== undefined && (obj.data = message.data ? Meal.toJSON(message.data) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<FindOneResponse>, I>>(base?: I): FindOneResponse {
    return FindOneResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<FindOneResponse>, I>>(object: I): FindOneResponse {
    const message = createBaseFindOneResponse();
    message.status = object.status ?? 0;
    message.error = object.error ?? "";
    message.data = (object.data !== undefined && object.data !== null) ? Meal.fromPartial(object.data) : undefined;
    return message;
  },
};

export interface MealService {
  CreateMeal(request: CreateMealRequest): Promise<CreateMealResponse>;
  FindOne(request: FindOneRequest): Promise<FindOneResponse>;
}

export class MealServiceClientImpl implements MealService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "meal.MealService";
    this.rpc = rpc;
    this.CreateMeal = this.CreateMeal.bind(this);
    this.FindOne = this.FindOne.bind(this);
  }
  CreateMeal(request: CreateMealRequest): Promise<CreateMealResponse> {
    const data = CreateMealRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateMeal", data);
    return promise.then((data) => CreateMealResponse.decode(_m0.Reader.create(data)));
  }

  FindOne(request: FindOneRequest): Promise<FindOneResponse> {
    const data = FindOneRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "FindOne", data);
    return promise.then((data) => FindOneResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
