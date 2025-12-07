#include <stdio.h>
#include <string.h>

enum commands{
    INSTALL,
    REMOVE,
    CommandCount
};

//gets the base command, which determines where the program goes next
int get_base_command(char* names[CommandCount], char **args)
{
    char* baseCommand = args[1];

    int command = -1;

    for(int i = 0; i < CommandCount; i++)
    {
        if (strcmp(baseCommand, names[i]) == 0)
        {
            command = i;
        }
    }

    return command;
}
#pragma region base commands
void install(char **args)
{
    printf("running install...");
}
#pragma endregion base commands

int main(int argc, char **argv)
{
    if (argc == 1)
    {
        char* message = "Display commands that can be used here \n"
                        "Yeah legit I got nothing here yet sorry";

        printf(message);

        return 10;
    }

    char* commandNames[CommandCount];

    commandNames[INSTALL] = "install";

    enum commands baseCommand = get_base_command(commandNames, argv);

    if (baseCommand == -1)
    {
        printf("The %s command doesn't exist", argv[1]);

        return 10;
    }

    if (baseCommand == INSTALL)
    {
        install(argv);
    }

    return 0;
}
