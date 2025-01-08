import * as jspb from 'google-protobuf'



export class Endpoint extends jspb.Message {
  getId(): string;
  setId(value: string): Endpoint;

  getProjectId(): string;
  setProjectId(value: string): Endpoint;

  getMethod(): string;
  setMethod(value: string): Endpoint;

  getPath(): string;
  setPath(value: string): Endpoint;

  getSetting(): Setting | undefined;
  setSetting(value?: Setting): Endpoint;
  hasSetting(): boolean;
  clearSetting(): Endpoint;

  getWorkflow(): Workflow | undefined;
  setWorkflow(value?: Workflow): Endpoint;
  hasWorkflow(): boolean;
  clearWorkflow(): Endpoint;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Endpoint.AsObject;
  static toObject(includeInstance: boolean, msg: Endpoint): Endpoint.AsObject;
  static serializeBinaryToWriter(message: Endpoint, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Endpoint;
  static deserializeBinaryFromReader(message: Endpoint, reader: jspb.BinaryReader): Endpoint;
}

export namespace Endpoint {
  export type AsObject = {
    id: string,
    projectId: string,
    method: string,
    path: string,
    setting?: Setting.AsObject,
    workflow?: Workflow.AsObject,
  }
}

export class Variable extends jspb.Message {
  getType(): VariableType;
  setType(value: VariableType): Variable;

  getRequired(): boolean;
  setRequired(value: boolean): Variable;

  getValue(): string;
  setValue(value: string): Variable;

  getDefault(): string;
  setDefault(value: string): Variable;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Variable.AsObject;
  static toObject(includeInstance: boolean, msg: Variable): Variable.AsObject;
  static serializeBinaryToWriter(message: Variable, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Variable;
  static deserializeBinaryFromReader(message: Variable, reader: jspb.BinaryReader): Variable;
}

export namespace Variable {
  export type AsObject = {
    type: VariableType,
    required: boolean,
    value: string,
    pb_default: string,
  }
}

export class Setting extends jspb.Message {
  getName(): string;
  setName(value: string): Setting;

  getDescription(): string;
  setDescription(value: string): Setting;

  getTimeoutMs(): number;
  setTimeoutMs(value: number): Setting;

  getNumWorkers(): number;
  setNumWorkers(value: number): Setting;

  getRequest(): SettingRequest | undefined;
  setRequest(value?: SettingRequest): Setting;
  hasRequest(): boolean;
  clearRequest(): Setting;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Setting.AsObject;
  static toObject(includeInstance: boolean, msg: Setting): Setting.AsObject;
  static serializeBinaryToWriter(message: Setting, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Setting;
  static deserializeBinaryFromReader(message: Setting, reader: jspb.BinaryReader): Setting;
}

export namespace Setting {
  export type AsObject = {
    name: string,
    description: string,
    timeoutMs: number,
    numWorkers: number,
    request?: SettingRequest.AsObject,
  }
}

export class SettingRequest extends jspb.Message {
  getQueryMap(): jspb.Map<string, Variable>;
  clearQueryMap(): SettingRequest;

  getJsonMap(): jspb.Map<string, Variable>;
  clearJsonMap(): SettingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SettingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SettingRequest): SettingRequest.AsObject;
  static serializeBinaryToWriter(message: SettingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SettingRequest;
  static deserializeBinaryFromReader(message: SettingRequest, reader: jspb.BinaryReader): SettingRequest;
}

export namespace SettingRequest {
  export type AsObject = {
    queryMap: Array<[string, Variable.AsObject]>,
    jsonMap: Array<[string, Variable.AsObject]>,
  }
}

export class Workflow extends jspb.Message {
  getStepsList(): Array<Step>;
  setStepsList(value: Array<Step>): Workflow;
  clearStepsList(): Workflow;
  addSteps(value?: Step, index?: number): Step;

  getEdgesList(): Array<Edge>;
  setEdgesList(value: Array<Edge>): Workflow;
  clearEdgesList(): Workflow;
  addEdges(value?: Edge, index?: number): Edge;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Workflow.AsObject;
  static toObject(includeInstance: boolean, msg: Workflow): Workflow.AsObject;
  static serializeBinaryToWriter(message: Workflow, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Workflow;
  static deserializeBinaryFromReader(message: Workflow, reader: jspb.BinaryReader): Workflow;
}

export namespace Workflow {
  export type AsObject = {
    stepsList: Array<Step.AsObject>,
    edgesList: Array<Edge.AsObject>,
  }
}

export class Step extends jspb.Message {
  getId(): string;
  setId(value: string): Step;

  getName(): string;
  setName(value: string): Step;

  getType(): StepType;
  setType(value: StepType): Step;

  getVariablesMap(): jspb.Map<string, Variable>;
  clearVariablesMap(): Step;

  getAction(): Action | undefined;
  setAction(value?: Action): Step;
  hasAction(): boolean;
  clearAction(): Step;

  getOutputsMap(): jspb.Map<string, Variable>;
  clearOutputsMap(): Step;

  getReturnsList(): Array<Return>;
  setReturnsList(value: Array<Return>): Step;
  clearReturnsList(): Step;
  addReturns(value?: Return, index?: number): Return;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Step.AsObject;
  static toObject(includeInstance: boolean, msg: Step): Step.AsObject;
  static serializeBinaryToWriter(message: Step, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Step;
  static deserializeBinaryFromReader(message: Step, reader: jspb.BinaryReader): Step;
}

export namespace Step {
  export type AsObject = {
    id: string,
    name: string,
    type: StepType,
    variablesMap: Array<[string, Variable.AsObject]>,
    action?: Action.AsObject,
    outputsMap: Array<[string, Variable.AsObject]>,
    returnsList: Array<Return.AsObject>,
  }
}

export class Action extends jspb.Message {
  getDataSourceId(): string;
  setDataSourceId(value: string): Action;

  getEnd(): ActionEnd | undefined;
  setEnd(value?: ActionEnd): Action;
  hasEnd(): boolean;
  clearEnd(): Action;

  getMysql(): ActionMysql | undefined;
  setMysql(value?: ActionMysql): Action;
  hasMysql(): boolean;
  clearMysql(): Action;

  getRest(): ActionRest | undefined;
  setRest(value?: ActionRest): Action;
  hasRest(): boolean;
  clearRest(): Action;

  getSleep(): ActionSleep | undefined;
  setSleep(value?: ActionSleep): Action;
  hasSleep(): boolean;
  clearSleep(): Action;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Action.AsObject;
  static toObject(includeInstance: boolean, msg: Action): Action.AsObject;
  static serializeBinaryToWriter(message: Action, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Action;
  static deserializeBinaryFromReader(message: Action, reader: jspb.BinaryReader): Action;
}

export namespace Action {
  export type AsObject = {
    dataSourceId: string,
    end?: ActionEnd.AsObject,
    mysql?: ActionMysql.AsObject,
    rest?: ActionRest.AsObject,
    sleep?: ActionSleep.AsObject,
  }
}

export class ActionEnd extends jspb.Message {
  getReturnDataFromStepIdsList(): Array<string>;
  setReturnDataFromStepIdsList(value: Array<string>): ActionEnd;
  clearReturnDataFromStepIdsList(): ActionEnd;
  addReturnDataFromStepIds(value: string, index?: number): ActionEnd;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ActionEnd.AsObject;
  static toObject(includeInstance: boolean, msg: ActionEnd): ActionEnd.AsObject;
  static serializeBinaryToWriter(message: ActionEnd, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ActionEnd;
  static deserializeBinaryFromReader(message: ActionEnd, reader: jspb.BinaryReader): ActionEnd;
}

export namespace ActionEnd {
  export type AsObject = {
    returnDataFromStepIdsList: Array<string>,
  }
}

export class ActionMysql extends jspb.Message {
  getQueriesList(): Array<Query>;
  setQueriesList(value: Array<Query>): ActionMysql;
  clearQueriesList(): ActionMysql;
  addQueries(value?: Query, index?: number): Query;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ActionMysql.AsObject;
  static toObject(includeInstance: boolean, msg: ActionMysql): ActionMysql.AsObject;
  static serializeBinaryToWriter(message: ActionMysql, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ActionMysql;
  static deserializeBinaryFromReader(message: ActionMysql, reader: jspb.BinaryReader): ActionMysql;
}

export namespace ActionMysql {
  export type AsObject = {
    queriesList: Array<Query.AsObject>,
  }
}

export class ActionRest extends jspb.Message {
  getPath(): Variable | undefined;
  setPath(value?: Variable): ActionRest;
  hasPath(): boolean;
  clearPath(): ActionRest;

  getMethod(): string;
  setMethod(value: string): ActionRest;

  getHeadersMap(): jspb.Map<string, Variable>;
  clearHeadersMap(): ActionRest;

  getRequestType(): string;
  setRequestType(value: string): ActionRest;

  getRequestBody(): string;
  setRequestBody(value: string): ActionRest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ActionRest.AsObject;
  static toObject(includeInstance: boolean, msg: ActionRest): ActionRest.AsObject;
  static serializeBinaryToWriter(message: ActionRest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ActionRest;
  static deserializeBinaryFromReader(message: ActionRest, reader: jspb.BinaryReader): ActionRest;
}

export namespace ActionRest {
  export type AsObject = {
    path?: Variable.AsObject,
    method: string,
    headersMap: Array<[string, Variable.AsObject]>,
    requestType: string,
    requestBody: string,
  }
}

export class ActionSleep extends jspb.Message {
  getTimeoutMs(): number;
  setTimeoutMs(value: number): ActionSleep;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ActionSleep.AsObject;
  static toObject(includeInstance: boolean, msg: ActionSleep): ActionSleep.AsObject;
  static serializeBinaryToWriter(message: ActionSleep, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ActionSleep;
  static deserializeBinaryFromReader(message: ActionSleep, reader: jspb.BinaryReader): ActionSleep;
}

export namespace ActionSleep {
  export type AsObject = {
    timeoutMs: number,
  }
}

export class Query extends jspb.Message {
  getQuery(): Variable | undefined;
  setQuery(value?: Variable): Query;
  hasQuery(): boolean;
  clearQuery(): Query;

  getParametersList(): Array<Variable>;
  setParametersList(value: Array<Variable>): Query;
  clearParametersList(): Query;
  addParameters(value?: Variable, index?: number): Variable;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Query.AsObject;
  static toObject(includeInstance: boolean, msg: Query): Query.AsObject;
  static serializeBinaryToWriter(message: Query, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Query;
  static deserializeBinaryFromReader(message: Query, reader: jspb.BinaryReader): Query;
}

export namespace Query {
  export type AsObject = {
    query?: Variable.AsObject,
    parametersList: Array<Variable.AsObject>,
  }
}

export class Return extends jspb.Message {
  getId(): string;
  setId(value: string): Return;

  getName(): string;
  setName(value: string): Return;

  getIsFinishCondition(): string;
  setIsFinishCondition(value: string): Return;

  getNextStepId(): string;
  setNextStepId(value: string): Return;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Return.AsObject;
  static toObject(includeInstance: boolean, msg: Return): Return.AsObject;
  static serializeBinaryToWriter(message: Return, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Return;
  static deserializeBinaryFromReader(message: Return, reader: jspb.BinaryReader): Return;
}

export namespace Return {
  export type AsObject = {
    id: string,
    name: string,
    isFinishCondition: string,
    nextStepId: string,
  }
}

export class Edge extends jspb.Message {
  getId(): string;
  setId(value: string): Edge;

  getConditionId(): string;
  setConditionId(value: string): Edge;

  getSource(): string;
  setSource(value: string): Edge;

  getDest(): string;
  setDest(value: string): Edge;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Edge.AsObject;
  static toObject(includeInstance: boolean, msg: Edge): Edge.AsObject;
  static serializeBinaryToWriter(message: Edge, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Edge;
  static deserializeBinaryFromReader(message: Edge, reader: jspb.BinaryReader): Edge;
}

export namespace Edge {
  export type AsObject = {
    id: string,
    conditionId: string,
    source: string,
    dest: string,
  }
}

export enum VariableType { 
  VARIABLE_TYPE_UNSPECIFIED = 0,
  VARIABLE_TYPE_STRING = 1,
  VARIABLE_TYPE_INT = 2,
  VARIABLE_TYPE_FLOAT = 3,
  VARIABLE_TYPE_BOOL = 4,
  VARIABLE_TYPE_OBJECT = 5,
}
export enum StepType { 
  STEP_TYPE_UNSPECIFIED = 0,
  STEP_TYPE_START = 1,
  STEP_TYPE_END = 2,
  STEP_TYPE_SLEEP = 3,
  STEP_TYPE_SCRIPT_JS = 4,
  STEP_TYPE_CONDITION = 5,
  STEP_TYPE_REST = 6,
  STEP_TYPE_MYSQL = 7,
  STEP_TYPE_POSTGRESQL = 8,
  STEP_TYPE_REDIS = 9,
}
