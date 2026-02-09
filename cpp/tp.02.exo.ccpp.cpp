#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* ChargerChaine(int N) {
    char* chaine = (char*)malloc((N + 1) * sizeof(char));
    if (chaine == NULL) {
        printf("Erreur de mamori.\n");
        exit(1);
    }
    printf("Veuillez entrer une chaine de caracteres (max %d): ", N);
    scanf(" %s", chaine); 
    return chaine;
}

int Longueur(char* ch) {
    int longueur = 0;
    while (ch[longueur] != '\0') {
        longueur++;
    }
    return longueur;
}

void ChargerTab(char Tab[], char* ch) {
    int i = 0;
    while (ch[i] != '\0') {
        Tab[i] = ch[i];
        i++;
    }
    Tab[i] = '\0';
}

void InverserTab(char Tab[], char T[], int m) {
    for (int i = 0; i < m; i++) {
        T[i] = Tab[m - i - 1];
    }
    T[m] = '\0';
}

void AfficherTab(char Tab[], int m) {
    for (int i = 0; i < m; i++) {
        printf("%c", Tab[i]);
    }
    printf("\n");
}

int main() {
    int N;
    printf("\nVeuillez la taille maximale de la chaine:\n ");
    scanf("%d", &N);

    char* ch = ChargerChaine(N);
    int m = Longueur(ch);

    char* Tab = (char*)malloc((m + 1) * sizeof(char));
    char* T = (char*)malloc((m + 1) * sizeof(char));
    if (Tab == NULL || T == NULL) {
        printf("\nErreur de mamori .");
        free(ch);
        exit(1);
    }

    ChargerTab(Tab, ch);
    InverserTab(Tab, T, m);

    printf("\nTableau original : ");
    AfficherTab(Tab, m);
    printf("\nTableau inverse : ");
    AfficherTab(T, m);

    free(ch);
    free(Tab);
    free(T);
    return 0;
}
