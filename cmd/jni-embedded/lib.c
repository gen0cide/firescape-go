#include <stdio.h>
#include <stdlib.h>
#include <jni.h>
#include "lib.h"


char * FirescapeInit(char *msg)
{

  JavaVMOption jvmopt[1];
  jvmopt[0].optionString = msg;

  JavaVMInitArgs vmArgs;
  vmArgs.version = JNI_VERSION_1_2;
  vmArgs.nOptions = 1;
  vmArgs.options = jvmopt;
  vmArgs.ignoreUnrecognized = JNI_TRUE;

  // Create the JVM
  JavaVM *javaVM;
  JNIEnv *jniEnv;
  long flag = JNI_CreateJavaVM(&javaVM, (void**)
    &jniEnv, &vmArgs);
  if (flag == JNI_ERR) {
    printf("Error creating VM. Exiting...\n");
    return msg;
  } 

  jclass jcls = (*jniEnv)->FindClass(jniEnv, "org/firescape/server/Server");
  if (jcls == NULL) {
    (*jniEnv)->ExceptionDescribe(jniEnv);
    (*javaVM)->DestroyJavaVM(javaVM);
    return msg;
  }
  if (jcls != NULL) {
    jmethodID methodId = (*jniEnv)->GetStaticMethodID(jniEnv, jcls,
      "EntryPoint", "(Ljava/lang/String;)V");
    if (methodId != NULL) {
      jstring str = (*jniEnv)->NewStringUTF(jniEnv, "jason");
      (*jniEnv)->CallStaticVoidMethod(jniEnv, jcls, methodId, str);
      if ((*jniEnv)->ExceptionCheck(jniEnv)) {
        (*jniEnv)->ExceptionDescribe(jniEnv);
        (*jniEnv)->ExceptionClear(jniEnv);
      }
    }
  }

  (*javaVM)->DestroyJavaVM(javaVM);
  return msg;
  }