import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import { EntGender } from './../../api/models/EntGender';
import { EntProvince } from './../../api/models/EntProvince';
import { EntUserType } from './../../api/models/EntUserType';


const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
     margin: theme.spacing(3),
   },
   withoutLabel: {
     marginTop: theme.spacing(3),
   },
   textField: {
     width: '25ch',
   },
 }),
);

export default function Create() {
 const classes = useStyles();
 const profile = { givenName: 'ระบบรับข้อมูลของ Guest เพื่อเป็น User' };
 const api = new DefaultApi();
 
 //const [user, setUser] = useState(initialUserState);
 const [genders, setGenders] = useState<EntGender[]>([]);
 const [provinces, setProvinces] = useState<EntProvince[]>([]);
 
 const [usertypes, setUserTypes] = useState<EntUserType[]>([]);
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 const [loading, setLoading] = useState(true);

 //const [user, setUser] = useState(String);
 

 const [usertypeids, setUserTypeid] = useState(Number);
 const [provinceids, setProvinceid] = useState(Number);
 const [genderids, setGenderid] = useState(Number);

 const [identitycard, setIdentityCard] = useState(Number);
 const [password, setPassword] = useState(String);
 const [confirmpassword, setConfirmPassword] = useState(String);
 const [firstname, setFirstName] = useState(String);
 const [lastname, setLastName] = useState(String);
 const [email, setEmail] = useState(String);
 const [phone, setPhone] = useState(Number);
 const [date, setDate] = useState(Number);

 useEffect(() => {
   const getGenders = async () => {
     const g = await api.listGender({ limit: 2, offset: 0 });
     setLoading(false);
     setGenders(g);
   };
   getGenders();

   const getProvinces = async () => {
    const pr = await api.listProvince({ limit: 7, offset: 0 });
    setLoading(false);
    setProvinces(pr);
  };
  getProvinces();  

  const getUserTypes = async () => {
    const ut = await api.listUsertype({ limit: 2, offset: 0 });
    setLoading(false);
    setUserTypes(ut);
  };
  getUserTypes();

 }, [loading]);
 
 const handleIdentityCardChange = (event: any) => {
  setIdentityCard(event.target.value as number);
 };

 const handlePasswordChange = (event: any) => {
  setPassword(event.target.value as string);

  
 };

 const handleConfirmPasswordChange = (event: any) => {
  setConfirmPassword(event.target.value as string);
 };

 const handleFirstNameChange = (event: any) => {
  setFirstName(event.target.value as string);
 };

 const handleLastNameChange = (event: any) => {
  setLastName(event.target.value as string);
 };

 const handleEmailChange = (event: any) => {
  setEmail(event.target.value as string);
 };

 const handlePhoneChange = (event: any) => {
  setPhone(event.target.value as number);
 };

 const handleDateChange = (event: any) => {
  setDate(event.target.value as number);
 };
   
 const CreateUser = async () =>{
  const user = {
    userType    : usertypeids,
    identityCard : +identitycard,
    password : password,
    confirmPassword : confirmpassword,
    firstName : firstname,
    lastName : lastname,
    gender : genderids,
    email : email,
    province : provinceids,
    phone : +phone,
    dateOfBirth : +date
  }
  
  console.log(user)
   const res:any = await api.createUser({ user : user });
   setStatus(true);
   if (res.id != ''){
     setAlert(true);
   } else {
     setAlert(false);
   }
   const timer = setTimeout(() => {
     setStatus(false);
   }, 1000);
 };


 const gender_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setGenderid(event.target.value as number);
 };

const province_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setProvinceid(event.target.value as number);
 };

 const usertype_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setUserTypeid(event.target.value as number);
 };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={` ${profile.givenName || 'ระบบประว้ติ'}`}
     ></Header>
     <Content>
       <ContentHeader title="ระบบสมัคร">
      
       {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                ✔️ ADD DATA COMPLETE : 😁 😆 🤩!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 ❌ CAN'T ADD DATA : 🥺 😵 😱!
               </Alert>
             )}
           </div>
         ) : null}
       </ContentHeader>





       <div className={classes.root}>
         <form noValidate autoComplete="off">


         
         <div>

<FormControl 
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="usertype-label">usertype</InputLabel>
            <Select
              labelId="usertype-label"
              label="userType"
              id="usertype_id"
              value={usertypeids}
              onChange={usertype_id_handleChange}
              style = {{width: 600}}
            >
              {usertypes.map((item:EntUserType)=>
              <MenuItem value={item.id}>{item.userTypeName}</MenuItem>)}
            </Select>
          </FormControl>   
 </div>

 <div>
 <FormControl
            // fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField
               id="identity"
               label="identityCard"
               variant="outlined"
               type="int"
               size="medium"
               value={identitycard}
               onChange={handleIdentityCardChange}
               style = {{width: 600}}
             />
           </FormControl>
           </div>

         <div>

           
         <FormControl
            // fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField
               id="password"
               label="Password"
               variant="outlined"
               type="password"
               size="medium"
               value={password}
               onChange={handlePasswordChange}
               style = {{width: 600}}
             />
           </FormControl>
           </div>        
<div>
           <FormControl
            // fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField
               id="confirmpassword"
               label="ConfirmPassword"
               variant="outlined"
               type="password"
               size="medium"
               value={confirmpassword}
               onChange={handleConfirmPasswordChange}
               style = {{width: 600}}
             />
           </FormControl>

           </div>



           <div>
           <FormControl
            // fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField
               id="firstname"
               label="FirstName"
               variant="outlined"
               type="string"
               size="medium"
               value={firstname}
               onChange={handleFirstNameChange}
               style = {{width: 600}}
             />
           </FormControl>
           </div>
<div>

           <FormControl
            // fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField
               id="lastname"
               label="LastName"
               variant="outlined"
               type="string"
               size="medium"
               value={lastname}
               onChange={handleLastNameChange}
               style = {{width: 600}}
             />
           </FormControl>
           </div>




           <div>

<FormControl 
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="gender-label">gender</InputLabel>
            <Select
              labelId="gender-label"
              label="Gender"
              id="gender_id"
              value={genderids}
              onChange={gender_id_handleChange}
              style = {{width: 600}}
            >
              {genders.map((item:EntGender)=>
              <MenuItem value={item.id}>{item.genderName}</MenuItem>)}
            </Select>
          </FormControl>   
 </div>




 
 <div>


 <FormControl
            // fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField
               id="email"
               label="Email"
               variant="outlined"
               type="string"
               size="medium"
               value={email}
               onChange={handleEmailChange}
               style = {{width: 600}}
             />
           </FormControl>
           </div>    



           <div>

<FormControl 
             className={classes.margin}
             variant="outlined"
           >
            <InputLabel id="province-label">province</InputLabel>
            <Select
              labelId="province-label"
              label="Province"
              id="provinceID"
              value={provinceids}
              onChange={province_id_handleChange}
              style = {{width: 600}}
            >
              {provinces.map((item:EntProvince)=>
              <MenuItem value={item.id}>{item.provinceName}</MenuItem>)}
            </Select>
          </FormControl>   
 </div>



<div>
 <FormControl
            // fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField
               id="phone"
               label="Phone"
               variant="outlined"
               type="int"
               size="medium"
               value={phone}
               onChange={handlePhoneChange}
               style = {{width: 600}}
             />
           </FormControl>
           </div>



           <div>

           <FormControl
            // fullWidth
             className={classes.margin}
             variant="outlined"
           >
             <TextField
               id="date"
               label="Date"
               variant="outlined"
               type="int"
               size="medium"
               value={date}
               onChange={handleDateChange}
               style = {{width: 600}}
             />
           </FormControl>


           </div>

           
           <div className={classes.margin}>
             <Button
               onClick={() => {
                 CreateUser();
               }}
               variant="contained"
               color="primary"
             >
               Submit
             </Button>
             
             <Button
               style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/"
               variant="contained"
               color="secondary"
             >
               Back
             </Button>
           </div>


           
         </form>
       </div>
     </Content>
   </Page>
 );
}