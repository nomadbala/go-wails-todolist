// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {db} from '../models';

export function CreateTask(arg1:string,arg2:boolean,arg3:string,arg4:string):Promise<void>;

export function DeleteTask(arg1:number):Promise<void>;

export function GetAllTasks():Promise<Array<db.Task>>;

export function Greet(arg1:string):Promise<string>;

export function MarkTaskDone(arg1:number):Promise<void>;

export function UpdateTask(arg1:number,arg2:string,arg3:boolean,arg4:string,arg5:string):Promise<void>;
