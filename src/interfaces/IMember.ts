import { EmployeeInterface } from "./IEmployee";
import { ProvinceInterface } from "./IProvince";
import { GendersInterface } from "./IGender";

export interface MemberInterface {
    ID?: number,
    FirstName?: string,
	LastName?:  string,
	Age?:       number,

	GenderID?: number,
	Gender?: GendersInterface;

	Date_Of_Birth?: Date | null;	


	ProvinceID?:number,
	Province?:ProvinceInterface,

	Telephone?: string,
	
	EmployeeID?: string,
	Employee?:EmployeeInterface,

}