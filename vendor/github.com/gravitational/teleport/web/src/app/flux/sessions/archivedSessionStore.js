/*
Copyright 2015 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import { Store, toImmutable } from 'nuclear-js';
import { Record, List } from 'immutable';
import { RECEIVE_SITE_EVENTS } from './actionTypes';
import { EventTypeEnum } from 'app/lib/term/enums';

const StoredSessionRec = Record({
  id: undefined,
  user: undefined,
  created: undefined,
  nodeIp: undefined,  
  last_active: undefined,
  server_id: undefined,
  siteId: undefined,
  parties: List()
}) 
    
export default Store({
  getInitialState() {
    return toImmutable({});
  },

  initialize() {        
    this.on(RECEIVE_SITE_EVENTS, receive);        
  }
})

// uses events to build stored session objects
function receive(state, { siteId, json }) {  
  const jsonEvents = json || [];
  let tmp = {};    
  return state.withMutations(state => {    
    jsonEvents.forEach( item => {
      if(item.event !== EventTypeEnum.START && item.event !== EventTypeEnum.END){
        return;
      }

      const { sid, user, time, event, server_id } = item;
      
      tmp[sid] = tmp[sid] || {}              
      tmp[sid].id = sid;
      tmp[sid].user = user;
      tmp[sid].siteId = siteId;

      if(event === EventTypeEnum.START){
        tmp[sid].created = time;
        tmp[sid].server_id = server_id;
        tmp[sid].nodeIp = item['addr.local'];
        tmp[sid].parties = [{
          user: user,
          serverIp: item['addr.remote']          
        }]
      }

      // update the store only with new items
      if(event === EventTypeEnum.END && !state.has(sid)){
        tmp[sid].last_active = time;
        state.set(sid, new StoredSessionRec(toImmutable(tmp[sid])));        
      }                  
    })    
  });  
}
