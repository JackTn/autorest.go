/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { clientAdapter } from './clients.js';
import { typeAdapter } from './types.js';
import { GoEmitterOptions } from '../lib.js';
import * as go from '../../../codemodel.go/src/gocodemodel.js';
import { packageNameFromOutputFolder, trimPackagePrefix } from '../../../naming.go/src/naming.js';
import * as tcgc from '@azure-tools/typespec-client-generator-core';
import { EmitContext } from '@typespec/compiler';
import { values } from '@azure-tools/linq';

const headerText = `Copyright (c) Microsoft Corporation. All rights reserved.
Licensed under the MIT License. See License.txt in the project root for license information.
Code generated by Microsoft (R) Go Code Generator.`;

export function tcgcToGoCodeModel(context: EmitContext<GoEmitterOptions>): go.CodeModel {
  const info = new go.Info('TODO Title');
  const options = new go.Options(headerText, context.options['generate-fakes'] === true, context.options['inject-spans'] === true, context.options['disallow-unknown-fields'] === true);
  if (context.options['azcore-version']) {
    options.azcoreVersion = context.options['azcore-version'];
  }

  const sdkContext = tcgc.createSdkContext(context);
  let codeModelType: go.CodeModelType = 'data-plane';
  if (values(sdkContext.experimental_sdkPackage.clients).any(client => { return client.arm; })) {
    // if one client is ARM then mark the code model as ARM.
    // we should never have an SDK that mixes ARM and data-plane.
    // TODO: tcgc to move arm property to sdkPackage
    codeModelType = 'azure-arm';
  }

  const codeModel = new go.CodeModel(info, codeModelType, packageNameFromOutputFolder(context.emitterOutputDir), options);
  if (context.options.module && context.options['module-version']) {
    codeModel.options.module = new go.Module(context.options.module, context.options['module-version']);
  } else if (context.options.module || context.options['module-version']) {
    throw new Error('--module and --module-version must both or neither be set');
  }
  if (context.options['rawjson-as-bytes']) {
    codeModel.options.rawJSONAsBytes = true;
  }
  if (context.options['slice-elements-byval']) {
    codeModel.options.sliceElementsByval = true;
  }

  fixStutteringTypeNames(sdkContext.experimental_sdkPackage, codeModel);

  const ta = new typeAdapter(codeModel);
  ta.adaptTypes(sdkContext);

  const ca = new clientAdapter(ta, context.options);
  ca.adaptClients(sdkContext.experimental_sdkPackage);
  codeModel.sortContent();
  return codeModel;
}

function fixStutteringTypeNames(sdkPackage: tcgc.SdkPackage<tcgc.SdkHttpOperation>, codeModel: go.CodeModel): void {
  let stutteringPrefix = codeModel.packageName;

  // if there's a well-known prefix, remove it
  if (stutteringPrefix.startsWith('arm')) {
    stutteringPrefix = stutteringPrefix.substring(3);
  } else if (stutteringPrefix.startsWith('az')) {
    stutteringPrefix = stutteringPrefix.substring(2);
  }
  stutteringPrefix = stutteringPrefix.toUpperCase();

  // ensure that enum, client, and struct type names don't stutter

  for (const sdkClient of sdkPackage.clients) {
    sdkClient.name = trimPackagePrefix(stutteringPrefix, sdkClient.name);
  }

  // check if the name collides with an existing name. we only do
  // this for model types as clients and enums get a suffix.
  const nameCollision = function(newName: string): boolean {
    for (const modelType of sdkPackage.models) {
      if (modelType.name === newName) {
        return true;
      }
    }
    return false;
  };

  // tracks type name collilsions due to renaming
  const collisions = new Array<string>();

  // trims the stuttering prefix from typeName and returns the new name.
  // if there's a collision, an entry is added to the collision list.
  const renameType = function(typeName: string): string {
    const originalName = typeName;
    const newName = trimPackagePrefix(stutteringPrefix, originalName); 

    // if the type was renamed to remove stuttering, check if it collides with an existing type name
    if (newName !== originalName && nameCollision(newName)) {
      collisions.push(`type ${originalName} was renamed to ${newName} which collides with an existing type name`);
    }
    return newName;
  };

  for (const sdkEnum of sdkPackage.enums) {
    sdkEnum.name = renameType(sdkEnum.name);
  }

  for (const modelType of sdkPackage.models) {
    modelType.name = renameType(modelType.name);
  }

  if (collisions.length > 0) {
    throw new Error(collisions.join('\n'));
  }
}
