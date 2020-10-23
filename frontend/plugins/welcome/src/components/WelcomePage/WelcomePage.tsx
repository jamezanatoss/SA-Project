import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';

const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'ระบบรับข้อมูลของ Guest เพื่อเป็น User' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={` ${profile.givenName || 'ระบบรับข้อมูลของ Guest'}`}
      
     >

        

     </Header>
     
     <Content>
        <ContentHeader title="">
         <Link component={RouterLink} to="/users">
           <Button variant="outlined" color="primary">
             Add Data
           </Button>
           </Link> 
                
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;