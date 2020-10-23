/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntGender,
    EntGenderFromJSON,
    EntGenderFromJSONTyped,
    EntGenderToJSON,
    EntProvince,
    EntProvinceFromJSON,
    EntProvinceFromJSONTyped,
    EntProvinceToJSON,
    EntUserType,
    EntUserTypeFromJSON,
    EntUserTypeFromJSONTyped,
    EntUserTypeToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUserEdges
 */
export interface EntUserEdges {
    /**
     * 
     * @type {EntGender}
     * @memberof EntUserEdges
     */
    genderID?: EntGender;
    /**
     * 
     * @type {EntProvince}
     * @memberof EntUserEdges
     */
    provinceID?: EntProvince;
    /**
     * 
     * @type {EntUserType}
     * @memberof EntUserEdges
     */
    userTypeID?: EntUserType;
}

export function EntUserEdgesFromJSON(json: any): EntUserEdges {
    return EntUserEdgesFromJSONTyped(json, false);
}

export function EntUserEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUserEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'genderID': !exists(json, 'GenderID') ? undefined : EntGenderFromJSON(json['GenderID']),
        'provinceID': !exists(json, 'ProvinceID') ? undefined : EntProvinceFromJSON(json['ProvinceID']),
        'userTypeID': !exists(json, 'UserTypeID') ? undefined : EntUserTypeFromJSON(json['UserTypeID']),
    };
}

export function EntUserEdgesToJSON(value?: EntUserEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'genderID': EntGenderToJSON(value.genderID),
        'provinceID': EntProvinceToJSON(value.provinceID),
        'userTypeID': EntUserTypeToJSON(value.userTypeID),
    };
}


