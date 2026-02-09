#include <stdio.h>

int main(){
	int num1,num2;
	
	//solicita ao usuario que insira dois numeros inteiros 
	printf("digite o primero numero:");
	scanf("%d",&num1);
	
	printf("digite o segundo numero:");
	scanf("%d",&num2);
	
	//verifique qual numero e maior 
	if(num1>num2){
	    printf("o maior numero e :%d\n",num1);
	} else if (num2>num1){
		printf("o maior numero e :%d\n",num2);
	}else{
		printf("os numeros sao iguais.\n");
	}
	return 0;
	
		}	
