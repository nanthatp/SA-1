import { EmployeeInterface } from "./IEmployee";
import { MemberInterface } from "./IMember";


export interface CartInterface {
  ID?: number;
  EmployeeID?: string;
  Employee?: EmployeeInterface;
  MemberID?: number;
  Member?: MemberInterface;
  Telephone?: string,
  Telephone_num?: MemberInterface,
}